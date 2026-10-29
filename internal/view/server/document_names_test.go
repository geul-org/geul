//ff:func feature=server type=helper control=sequence
//ff:what TestDocumentNames — names of that Issuer's documents.geul; other Issuers and missing files give none
package server

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDocumentNames(t *testing.T) {
	root := namesFixture(t)
	if n := documentNames(treeOf(root), "0x0000"); len(n) != 2 || n["sample-b"] != "0xffff000000000002" || n["publication"] != "0x0000000000001018" {
		t.Fatalf("got %v", n)
	}
	other := filepath.Join(root, ".well-known", "geul", "0x0001")
	os.MkdirAll(other, 0o755)
	b, _ := os.ReadFile(filepath.Join(root, ".well-known", "geul", "0x0000", "documents.geul"))
	os.WriteFile(filepath.Join(other, "documents.geul"), b, 0o644)
	if n := documentNames(treeOf(root), "0x0001"); len(n) != 0 {
		t.Fatalf("foreign: %v", n)
	}
	if n := documentNames(treeOf(root), "0x0002"); len(n) != 0 {
		t.Fatalf("missing: %v", n)
	}
}
