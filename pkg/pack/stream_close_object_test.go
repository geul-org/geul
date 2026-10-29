//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamCloseObject — the object is named by its SHA-256 under o/ with header, items and tail; a closed part file or a missing o/ fails
package pack

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"testing"
)

func TestStreamCloseObject(t *testing.T) {
	s := streamTestNew(t, nil)
	f, open, _ := s.startObject(kindData, streamDataPart)
	f.WriteAt([]byte("record"), int64(open.Off))
	open.Off, open.Count = open.Off+6, 1
	if err := s.closeObject(f, open, kindData, []byte("inv")); err != nil {
		t.Fatal(err)
	}
	o := s.st.Objects[0]
	b, err := os.ReadFile(filepath.Join(s.out, "o", hex.EncodeToString(o.Hash[:])+".gpk"))
	if err != nil || sha256.Sum256(b) != o.Hash || o.Size != headerSize+9 || string(b[headerSize:]) != "recordinv" {
		t.Fatalf("%+v %v", o, err)
	}
	f, open, _ = s.startObject(kindIndex, streamIndexPart)
	f.Close()
	if err := s.closeObject(f, open, kindIndex, nil); err == nil {
		t.Fatal("closed part: want error")
	}
	f, open, _ = s.startObject(kindIndex, streamIndexPart)
	os.RemoveAll(filepath.Join(s.out, "o"))
	if err := s.closeObject(f, open, kindIndex, nil); err == nil {
		t.Fatal("no o/: want error")
	}
}
