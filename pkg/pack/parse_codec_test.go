//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestParseCodec — "" follows (nil), none is Stored, zstd forms are Zstd; anything else is an error
package pack

import "testing"

func TestParseCodec(t *testing.T) {
	if c, err := ParseCodec(" "); c != nil || err != nil {
		t.Fatal(c, err)
	}
	if c, err := ParseCodec("None"); c != (Stored{}) || err != nil {
		t.Fatal(c, err)
	}
	if c, err := ParseCodec("zstd:7:nodict"); err != nil || c.(*Zstd).Level != 7 || c.(*Zstd).DictSize != -1 {
		t.Fatal(c, err)
	}
	for _, s := range []string{":", "none:1", "gzip", "zstd:0", "zstd:x"} {
		if _, err := ParseCodec(s); err == nil {
			t.Fatal(s)
		}
	}
}
