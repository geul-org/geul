//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamStoreDict — the dictionary is the first record (inventory kind 4) and the one records use; a stream that cannot place it fails
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStreamStoreDict(t *testing.T) {
	s := streamTestNew(t, nil)
	d := &Dict{ID: FirstDictID, Bytes: []byte("dictionary bytes")}
	if err := s.storeDict(d); err != nil || s.dict != d || s.st.Inv[0].Kind != invDict || s.st.Dict.Loc.Off != headerSize || s.st.Res.Dict != FirstDictID {
		t.Fatalf("%+v %v", s.st.Inv, err)
	}
	s2 := streamTestNew(t, nil)
	os.RemoveAll(filepath.Join(s2.out, streamDir))
	if err := s2.storeDict(d); err == nil {
		t.Fatal("no working folder: want error")
	}
}
