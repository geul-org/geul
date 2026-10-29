//ff:func feature=server type=helper control=sequence
//ff:what TestOwnNames — only the given Issuer's documents; broken bytes give none
package server

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOwnNames(t *testing.T) {
	b, _ := os.ReadFile(filepath.Join(namesFixture(t), ".well-known", "geul", "0x0000", "documents.geul"))
	if n := ownNames(b, "0x0000"); len(n) != 2 || n["sample-b"] != "0xffff000000000002" {
		t.Fatal(n)
	}
	if n := ownNames(b, "0x0001"); len(n) != 0 {
		t.Fatal(n)
	}
	if n := ownNames([]byte{1}, "0x0000"); len(n) != 0 {
		t.Fatal(n)
	}
}
