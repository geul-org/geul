//ff:func feature=cli type=helper control=sequence
//ff:what TestAppDir — the data folder as an archive; none set or found fails
package cli

import (
	"testing"

	"github.com/geul-org/geul/go/internal/config"
)

func TestAppDir(t *testing.T) {
	a := &app{env: config.Env{Getenv: func(string) string { return "" }, Cwd: t.TempDir(), ConfigDir: t.TempDir()}}
	a.cfg, _ = config.Load(a.env)
	if _, err := a.dir(); err == nil {
		t.Fatal("none: want error")
	}
	a.env.Flags = map[string]string{"GEULDATA": "/d"}
	a.cfg, _ = config.Load(a.env)
	if d, err := a.dir(); err != nil || d.Root != "/d" {
		t.Fatal(err)
	}
}
