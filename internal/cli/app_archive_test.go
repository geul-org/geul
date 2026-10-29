//ff:func feature=cli type=helper control=sequence
//ff:what TestAppArchive — opened once and kept; a directory archive is remembered as nil
package cli

import (
	"testing"

	"github.com/geul-org/gogeul/internal/config"
)

func TestAppArchive(t *testing.T) {
	h := newHarness(t)
	h.doc()
	a := &app{}
	a.cfg, _ = config.Load(config.Env{Getenv: func(string) string { return "" }, Cwd: t.TempDir(), ConfigDir: t.TempDir()})
	if arch, err := a.archive(h.data); arch != nil || err != nil || !a.dirs {
		t.Fatal(arch, err)
	}
	h.packed()
	if arch, _ := a.archive(h.data); arch != nil {
		t.Fatal("remembered as a directory")
	}
	a = &app{cfg: a.cfg}
	arch, err := a.archive(h.data)
	if err != nil || arch == nil {
		t.Fatal(err)
	}
	if again, _ := a.archive(h.data); again != arch {
		t.Fatal("not kept")
	}
}
