//ff:func feature=pack type=encoder control=sequence
//ff:what TestNewEncoder — no dictionary, raw content and formatted dictionaries name their ID in frames; an ID mismatch and a broken dictionary fail
package pack

import (
	"os"
	"strings"
	"testing"

	"github.com/klauspost/compress/zstd"
)

func TestNewEncoder(t *testing.T) {
	src := []byte(strings.Repeat("geul pack record ", 40))
	e, err := newEncoder(zstd.SpeedDefault, 1, nil)
	if h, _ := parseFrameHeader(e.EncodeAll(src, nil)); err != nil || h.dictID != 0 || !h.checksum || h.fcs != uint64(len(src)) {
		t.Fatal(err, h)
	}
	e, err = newEncoder(zstd.SpeedDefault, 1, &Dict{ID: 77, Bytes: []byte("geul pack record geul pack record")})
	if h, _ := parseFrameHeader(e.EncodeAll(src, nil)); err != nil || h.dictID != 77 {
		t.Fatal(err, h)
	}
	cli, _ := os.ReadFile("testdata/zstd/cli.dict")
	e, err = newEncoder(zstd.SpeedDefault, 1, &Dict{ID: FirstDictID, Bytes: cli})
	if h, _ := parseFrameHeader(e.EncodeAll(src, nil)); err != nil || h.dictID != FirstDictID {
		t.Fatal(err, h)
	}
	if _, err := newEncoder(zstd.SpeedDefault, 1, &Dict{ID: 1, Bytes: cli}); err == nil {
		t.Fatal("ID mismatch")
	}
	if _, err := newEncoder(zstd.SpeedDefault, 1, &Dict{ID: FirstDictID, Bytes: cli[:12]}); err == nil {
		t.Fatal("broken dictionary")
	}
}
