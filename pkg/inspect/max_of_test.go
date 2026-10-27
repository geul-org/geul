//ff:func feature=inspect type=helper control=sequence
//ff:what TestMaxOf — the largest of the first and the rest; no rest gives the first
package inspect

import "testing"

func TestMaxOf(t *testing.T) {
	if maxOf(3, 9, 1) != 9 || maxOf(5) != 5 {
		t.Fatal("maxOf")
	}
}
