//ff:func feature=cli type=helper control=sequence
//ff:what TestPackSources — packed file:// folders read through their archive, GEULPACK items first; bad items and broken archives fail
package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestPackSources(t *testing.T) {
	h := newHarness(t)
	h.doc()
	plain := h.data
	h.packed()
	cfg, _ := resolve.ConfigFromEnv(func(k string) string {
		if k == "GEULPROXY" {
			return "file://" + h.data + ",file://" + plain + ",registry"
		}
		return ""
	})
	set := filepath.Join(h.data, ".well-known", "geul", "0x0000")
	cfg, err := packSources(cfg, " "+set+" ,,")
	if err != nil || len(cfg.Proxy) != 4 || cfg.Proxy[0].Name != "pack+"+set {
		t.Fatal(cfg.Proxy, err)
	}
	if d := cfg.Proxy[1].Source.(resolve.Dir); d.Store == nil {
		t.Fatal("packed file:// folder")
	}
	if d := cfg.Proxy[2].Source.(resolve.Dir); d.Store != nil {
		t.Fatal("plain file:// folder")
	}
	if _, err := packSources(resolve.Config{}, t.TempDir()); err == nil {
		t.Fatal("bad GEULPACK item")
	}
	os.WriteFile(filepath.Join(set, "latest.txt"), []byte("x"), 0o644)
	cfg, _ = resolve.ConfigFromEnv(func(k string) string {
		if k == "GEULPROXY" {
			return "file://" + h.data
		}
		return ""
	})
	if _, err := packSources(cfg, ""); err == nil {
		t.Fatal("broken archive")
	}
}
