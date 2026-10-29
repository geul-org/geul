//ff:func feature=server type=helper control=sequence
//ff:what TestTreeLabels — Issuer names at the top, Issuer and document names under an Issuer, codec spec names under the GEUL Issuer
package server

import (
	"testing"
)

func TestTreeLabels(t *testing.T) {
	root := namesFixture(t)
	if l := treeLabels(treeOf(root), "/.well-known/geul/"); len(l) != 1 || l["0x0000"] != "geul.org" {
		t.Fatalf("top: %v", l)
	}
	if l := treeLabels(treeOf(root), "/.well-known/geul/0x0000/"); len(l) != 3 || l["0xffff000000000002"] != "sample-b" {
		t.Fatalf("issuer: %v", l)
	}
	if l := treeLabels(treeOf(root), "/.well-known/geul/0x0000/"); l["0x0000000000001018"] != "publication" {
		t.Fatalf("codec spec: %v", l)
	}
}
