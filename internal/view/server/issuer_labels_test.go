//ff:func feature=server type=helper control=sequence
//ff:what TestIssuerLabels — the inverted names of issuers.geul; none when the tree has no issuers.geul
package server

import "testing"

func TestIssuerLabels(t *testing.T) {
	if l := issuerLabels(namesFixture(t)); len(l) != 1 || l["0x0000"] != "geul.org" {
		t.Fatalf("got %v", l)
	}
	if l := issuerLabels(t.TempDir()); len(l) != 0 {
		t.Fatalf("none: %v", l)
	}
}
