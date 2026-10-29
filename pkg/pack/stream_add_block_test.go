//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestStreamAddBlock — blocks go to an index object, a full one is closed and the next started; a failing codec, a missing working folder and a closed part fail
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStreamAddBlock(t *testing.T) {
	s := streamTestNew(t, Stored{})
	var kvs []streamKV
	for i := 0; i < 40; i++ {
		kvs = append(kvs, streamKV{Key: pKey(0, uint64(i)*1000003, 1), Loc: Loc{Off: uint64(i) * 977, Len: 900}})
	}
	a, err := s.addBlock(blockOf(blkPLeaf, kvs))
	if err != nil || a.Obj != 0 || a.Off != headerSize || a.Enc != encZstd {
		t.Fatalf("%+v %v", a, err)
	}
	for i := 0; i < 6; i++ {
		a, err = s.addBlock(blockOf(blkPLeaf, kvs))
	}
	if err != nil || a.Obj == 0 || s.st.Objects[0].Size == 0 {
		t.Fatalf("expected a second index object: %+v %v", a, err)
	}
	s.index.Close()
	if _, err := s.addBlock(blockOf(blkPLeaf, kvs[:1])); err == nil {
		t.Fatal("closed part: want error")
	}
	if _, err := streamTestNew(t, failCodec{}).addBlock(blockOf(blkPLeaf, kvs[:1])); err == nil {
		t.Fatal("failing codec: want error")
	}
	s2 := streamTestNew(t, Stored{})
	os.RemoveAll(filepath.Join(s2.out, streamDir))
	if _, err := s2.addBlock(blockOf(blkPLeaf, kvs[:1])); err == nil {
		t.Fatal("no working folder: want error")
	}
	s3 := streamTestNew(t, Stored{})
	for i := 0; i < 3; i++ {
		s3.addBlock(blockOf(blkPLeaf, kvs))
	}
	os.RemoveAll(filepath.Join(s3.out, "o"))
	for i := 0; i < 6 && err == nil; i++ {
		_, err = s3.addBlock(blockOf(blkPLeaf, kvs))
	}
	if err == nil {
		t.Fatal("index object that cannot close: want error")
	}
}
