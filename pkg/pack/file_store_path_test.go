//ff:func feature=pack type=helper control=sequence
//ff:what TestFileStorePath — slash names under Root
package pack

import (
	"path/filepath"
	"testing"
)

func TestFileStorePath(t *testing.T) {
	if p := (FileStore{Root: "r"}).path("o/x.gpk"); p != filepath.Join("r", "o", "x.gpk") {
		t.Fatal(p)
	}
}
