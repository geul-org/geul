//ff:func feature=pack type=helper control=sequence
//ff:what TestZstdWorkers — 0 is DefaultWorkers; a set value stays
package pack

import "testing"

func TestZstdWorkers(t *testing.T) {
	if (&Zstd{}).workers() != DefaultWorkers || (&Zstd{Workers: 3}).workers() != 3 {
		t.Fatal("workers")
	}
}
