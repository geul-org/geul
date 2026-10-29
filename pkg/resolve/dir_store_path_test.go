//ff:func feature=resolve type=helper control=sequence
//ff:what TestDirStorePath — slash names under the root
package resolve

import (
	"path/filepath"
	"testing"
)

func TestDirStorePath(t *testing.T) {
	if p := (dirStore{root: "r"}).path("a/b.geul"); p != filepath.Join("r", "a", "b.geul") {
		t.Fatal(p)
	}
}
