//ff:func feature=pack type=helper control=sequence
//ff:what TestSortKVs — ascending keys
package pack

import "testing"

func TestSortKVs(t *testing.T) {
	if k := sortKVs([]kv{{"b", Loc{}}, {"a", Loc{}}}); k[0].key != "a" {
		t.Fatal(k)
	}
}
