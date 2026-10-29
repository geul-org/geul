//ff:func feature=pack type=decoder control=sequence
//ff:what TestNewDictDecoder — a zstd CLI dictionary and frame decode; raw content works; a wrong ID or a broken formatted dictionary is ErrCorrupt
package pack

import (
	"bytes"
	"errors"
	"os"
	"testing"
)

func TestNewDictDecoder(t *testing.T) {
	dict, _ := os.ReadFile("testdata/zstd/cli.dict")
	frame, _ := os.ReadFile("testdata/zstd/sample.txt.zst")
	want, _ := os.ReadFile("testdata/zstd/sample.txt")
	d, err := newDictDecoder(FirstDictID, dict)
	if err != nil {
		t.Fatal(err)
	}
	if b, err := d.DecodeAll(frame, nil); err != nil || !bytes.Equal(b, want) {
		t.Fatal(err)
	}
	if _, err := newDictDecoder(FirstDictID+1, dict); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
	if _, err := newDictDecoder(5, []byte("raw content")); err != nil {
		t.Fatal(err)
	}
	if _, err := newDictDecoder(FirstDictID, dict[:12]); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
