//ff:func feature=cli type=command control=iteration dimension=1
//ff:what app.envList — geul env: every setting as KEY=VALUE with its source, or as JSON
package cli

import (
	"fmt"

	"github.com/geul-org/gogeul/internal/config"
)

// envList prints each key of config.Keys and where its value came from.
func (a *app) envList() error {
	all := map[string]config.Value{}
	for _, k := range config.Keys {
		all[k] = config.Value{Value: a.cfg.Get(k), Source: a.cfg.Source(k)}
	}
	if done, err := a.json(all); done {
		return err
	}
	for _, k := range config.Keys {
		fmt.Fprintf(a.out, "%s=%s", k, a.cfg.Get(k))
		if s := a.cfg.Source(k); s != "" {
			fmt.Fprintf(a.out, "  # %s", s)
		}
		fmt.Fprintln(a.out)
	}
	return nil
}
