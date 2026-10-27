//ff:func feature=cli type=helper control=sequence
//ff:what harness.run — test fixture: run one geul command line with stdin; its output and error
package cli

import (
	"bytes"
	"strings"

	"github.com/geul-org/gogeul/internal/config"
)

// run builds a fresh command tree (newRootCmd) and executes args.
func (h *harness) run(stdin string, args ...string) (string, error) {
	var out bytes.Buffer
	a := &app{env: config.Env{Getenv: func(k string) string { return h.env[k] }, Cwd: h.cwd, ConfigDir: h.conf, CacheDir: h.conf}, in: strings.NewReader(stdin), out: &out}
	cmd := newRootCmd(a)
	cmd.SetArgs(args)
	err := cmd.Execute()
	return out.String(), err
}
