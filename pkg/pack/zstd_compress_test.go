//ff:func feature=pack type=encoder control=sequence
//ff:what TestZstdCompress — short input is a raw-block frame; longer input compresses with content size, checksum and the dictionary's ID, and decodes back; encoder errors fail
package pack

import (
	"bytes"
	"strings"
	"testing"
)

func TestZstdCompress(t *testing.T) {
	z := &Zstd{Workers: 1}
	if f, err := z.Compress([]byte("tiny"), nil); err != nil || !bytes.Equal(f, zstdFrame([]byte("tiny"))) {
		t.Fatal(err)
	}
	src := []byte(strings.Repeat("<record kind=\"ALIAS\" lang=\"en\"/>\n", 30))
	f, err := z.Compress(src, nil)
	h, _ := parseFrameHeader(f)
	if err != nil || len(f) >= len(src) || !h.hasFCS || h.fcs != uint64(len(src)) || !h.checksum || h.dictID != 0 {
		t.Fatal(err, h)
	}
	if b, err := plainDecoder().DecodeAll(f, nil); err != nil || !bytes.Equal(b, src) {
		t.Fatal(err)
	}
	d, _ := testDict(t, FirstDictID, 64)
	f, err = z.Compress(src, d)
	if h, _ := parseFrameHeader(f); err != nil || h.dictID != FirstDictID {
		t.Fatal(err, h)
	}
	dec, _ := newDictDecoder(d.ID, d.Bytes)
	if b, err := dec.DecodeAll(f, nil); err != nil || !bytes.Equal(b, src) {
		t.Fatal(err)
	}
	if _, err := z.Compress(src, &Dict{ID: 1, Bytes: d.Bytes}); err == nil {
		t.Fatal("dictionary ID mismatch")
	}
}
