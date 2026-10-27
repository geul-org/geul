//ff:func feature=server type=helper control=sequence
//ff:what TestNamePath — Issuer, document and codec spec names become IDs; a path with no name is false
package server

import (
	"testing"
)

func TestNamePath(t *testing.T) {
	root := namesFixture(t)
	if p, ok := namePath(root, "/.well-known/geul/geul.org/sample-b/x"); !ok || p != "/.well-known/geul/0x0000/0xffff000000000002/x" {
		t.Fatalf("both: %q %v", p, ok)
	}
	if p, ok := namePath(root, "/.well-known/geul/geul.org"); !ok || p != "/.well-known/geul/0x0000" {
		t.Fatalf("issuer: %q %v", p, ok)
	}
	if p, ok := namePath(root, "/.well-known/geul/geul.org/publication/"); !ok || p != "/.well-known/geul/0x0000/0x0000000000001018/" {
		t.Fatalf("codec: %q %v", p, ok)
	}
	if _, ok := namePath(root, "/.well-known/geul/0x0000/nope"); ok {
		t.Fatal("no name: want false")
	}
}
