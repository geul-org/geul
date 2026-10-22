//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestKindString — every kind has its spec name
package core

import "testing"

func TestKindString(t *testing.T) {
	want := map[Kind]string{KindReference: "REFERENCE", KindValue: "VALUE", KindEdge: "EDGE", KindMeta: "META"}
	for k, s := range want {
		if got := k.String(); got != s {
			t.Fatalf("%d: got %q want %q", k, got, s)
		}
	}
}
