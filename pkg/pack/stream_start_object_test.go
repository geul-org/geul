//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamStartObject — the next OBJECTS position is reserved and the part file is empty, items starting after the header; a missing working folder fails
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStreamStartObject(t *testing.T) {
	s := streamTestNew(t, nil)
	f, open, err := s.startObject(kindIndex, streamIndexPart)
	if err != nil || open.Num != 0 || open.Off != headerSize || len(s.st.Objects) != 1 || s.st.Objects[0].Kind != kindIndex {
		t.Fatalf("%+v %v", open, err)
	}
	if fi, _ := f.Stat(); fi.Size() != 0 {
		t.Fatalf("size %d: the header is written when the object closes", fi.Size())
	}
	os.RemoveAll(filepath.Join(s.out, streamDir))
	if _, _, err := s.startObject(kindData, streamDataPart); err == nil {
		t.Fatal("no working folder: want error")
	}
}
