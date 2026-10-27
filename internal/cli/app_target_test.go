//ff:func feature=cli type=helper control=sequence
//ff:what TestAppTarget — the given path wins; without one, the data folder, or an error when none is found
package cli

import (
	"testing"

	"github.com/geul-org/gogeul/internal/config"
)

func TestAppTarget(t *testing.T) {
	a := &app{cfg: config.Config{}}
	if p, err := a.target([]string{"x.geul"}); err != nil || p != "x.geul" {
		t.Fatalf("given: %q %v", p, err)
	}
	if _, err := a.target(nil); err == nil {
		t.Fatal("no data folder: want error")
	}
	h := newHarness(t)
	h.must("env")
	a = &app{env: config.Env{Getenv: func(string) string { return "" }, Cwd: h.data, ConfigDir: h.conf, CacheDir: h.conf}}
	a.cfg, _ = config.Load(a.env)
	if p, err := a.target(nil); err != nil || p != h.data {
		t.Fatalf("data folder: %q %v", p, err)
	}
}
