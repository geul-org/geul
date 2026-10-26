//ff:func feature=cli type=helper control=sequence
//ff:what harness.key — test fixture: resolve an address the way the commands do
package cli

import (
	"github.com/geul-org/geul/go/internal/config"
	"github.com/geul-org/geul/go/pkg/resolve"
)

// key loads the harness settings and resolves addr (app.key).
func (h *harness) key(addr string) (resolve.Key, error) {
	a := &app{env: config.Env{Getenv: func(k string) string { return h.env[k] }, Cwd: h.cwd, ConfigDir: h.conf}}
	a.cfg, _ = config.Load(a.env)
	return a.key(addr)
}
