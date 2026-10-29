//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirTree — the directory store under Root, or the given Store
package resolve

import (
	"testing"
	"testing/fstest"
)

func TestDirTree(t *testing.T) {
	if _, ok := (Dir{Root: "r"}).Tree().(dirStore); !ok {
		t.Fatal("dir")
	}
	if _, ok := (Dir{Store: mapStore{fstest.MapFS{}}}).Tree().(mapStore); !ok {
		t.Fatal("store")
	}
}
