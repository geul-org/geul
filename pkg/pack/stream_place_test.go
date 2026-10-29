//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamPlace — records fill a data pack, the next one past PackSize starts a new pack, an oversized record goes alone; failures to start, write or finish a pack are errors (closed part, missing folder, failing codec)
package pack

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestStreamPlace(t *testing.T) {
	s := streamTestNew(t, nil)
	a, err := s.place(bytes.Repeat([]byte{1}, 2000), encRaw)
	if err != nil || a.Obj != 0 || a.Off != headerSize {
		t.Fatalf("%+v %v", a, err)
	}
	b, err := s.place(bytes.Repeat([]byte{2}, 2000), encZstd)
	if err != nil || b.Obj != 1 || b.Off != headerSize || b.Enc != encZstd || len(s.st.Objects) != 2 || s.st.Objects[0].Size == 0 {
		t.Fatalf("%+v %v", b, err)
	}
	big, err := s.place(bytes.Repeat([]byte{3}, 5000), encRaw)
	if err != nil || big.Obj != 2 {
		t.Fatalf("oversized: %+v %v", big, err)
	}
	w := streamTestNew(t, nil)
	w.place([]byte{1}, encRaw)
	w.data.Close()
	if _, err := w.place([]byte{4}, encRaw); err == nil {
		t.Fatal("closed part: want error")
	}
	s.data.Close()
	if _, err := s.place([]byte{4}, encRaw); err == nil {
		t.Fatal("pack that cannot finish: want error")
	}
	s2 := streamTestNew(t, nil)
	os.RemoveAll(filepath.Join(s2.out, streamDir))
	if _, err := s2.place([]byte{1}, encRaw); err == nil {
		t.Fatal("no working folder: want error")
	}
	s3 := streamTestNew(t, failCodec{})
	s3.place(bytes.Repeat([]byte{1}, 2000), encRaw)
	if _, err := s3.place(bytes.Repeat([]byte{1}, 2000), encRaw); err == nil {
		t.Fatal("inventory that does not compress: want error")
	}
}
