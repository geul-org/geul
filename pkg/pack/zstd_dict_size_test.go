//ff:func feature=pack type=helper control=sequence
//ff:what TestZstdDictSize — 0 is DefaultDictSize; a set value stays
package pack

import "testing"

func TestZstdDictSize(t *testing.T) {
	if (&Zstd{}).dictSize() != DefaultDictSize || (&Zstd{DictSize: 4096}).dictSize() != 4096 {
		t.Fatal("dictSize")
	}
}
