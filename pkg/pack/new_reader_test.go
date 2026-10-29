//ff:func feature=pack type=resolver control=sequence
//ff:what TestNewReader — defaults: Stored codec and DefaultCacheBlocks; given values are kept
package pack

import "testing"

func TestNewReader(t *testing.T) {
	r := newReader(memStore{}, Options{}, &manifest{issuers: []string{"0x0000"}}, "n")
	if _, ok := r.codec.(Stored); !ok || r.limit != DefaultCacheBlocks || r.inum["0x0000"] != 0 || r.name != "n" {
		t.Fatal(r)
	}
	c := Stored{}
	r = newReader(memStore{}, Options{Codec: &c, CacheBlocks: 3}, &manifest{}, "n")
	if r.codec != &c || r.limit != 3 {
		t.Fatal(r)
	}
}
