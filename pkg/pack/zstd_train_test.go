//ff:func feature=pack type=encoder control=sequence
//ff:what TestZstdTrain — a formatted dictionary under the given ID, the same bytes every time, smaller frames than without it; no content fails
package pack

import (
	"bytes"
	"testing"
)

func TestZstdTrain(t *testing.T) {
	d, samples := testDict(t, FirstDictID+3, 128)
	if id, ok := dictID(d.Bytes); !ok || id != FirstDictID+3 || d.ID != id || len(d.Bytes) > 4096+1024 {
		t.Fatal(id, ok, len(d.Bytes))
	}
	again, _ := testDict(t, FirstDictID+3, 128)
	if !bytes.Equal(d.Bytes, again.Bytes) {
		t.Fatal("not deterministic")
	}
	z := &Zstd{Workers: 1}
	plain, _ := z.Compress(samples[5], nil)
	with, _ := z.Compress(samples[5], d)
	if len(with) >= len(plain) {
		t.Fatal(len(with), len(plain))
	}
	if _, err := (&Zstd{}).Train([][]byte{[]byte("x")}, FirstDictID); err == nil {
		t.Fatal("no content")
	}
}
