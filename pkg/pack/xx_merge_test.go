//ff:func feature=pack type=helper control=sequence
//ff:what TestXxMerge — merging a zero lane multiplies and adds the primes
package pack

import "testing"

func TestXxMerge(t *testing.T) {
	p1, p4 := xxP1, xxP4
	if xxMerge(1, 0) != p1+p4 {
		t.Fatal("merge")
	}
}
