//ff:func feature=config type=helper control=sequence
//ff:what TestDefaults — remote and format always; cache under the cache folder; the data folder only when found
package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaults(t *testing.T) {
	root := t.TempDir()
	os.MkdirAll(filepath.Join(root, ".well-known", "geul"), 0o755)
	d := defaults(root, "/c")
	if d["GEULDATA"].Value != root || d["GEULDATA"].Source != "found" || d["GEULCACHE"].Value != "/c/geul" || d["GEULFORMAT"].Value != "text" {
		t.Fatalf("%+v", d)
	}
	d = defaults(t.TempDir(), "")
	if _, ok := d["GEULCACHE"]; ok {
		t.Fatal("no cache folder: want no default")
	}
}
