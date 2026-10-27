//ff:func feature=cli type=command control=iteration dimension=1
//ff:what app.envWrite — geul env -w KEY=VALUE / -u KEY: change the settings file
package cli

import (
	"fmt"
	"strings"

	"github.com/geul-org/gogeul/internal/config"
)

// envWrite sets (set: each arg KEY=VALUE) or removes (each arg KEY) keys in
// the settings file in effect.
func (a *app) envWrite(args []string, set bool) error {
	for _, arg := range args {
		k, v, ok := strings.Cut(arg, "=")
		if set && (!ok || v == "") {
			return fmt.Errorf("env -w: %q is not KEY=VALUE", arg)
		}
		if err := config.Write(a.cfg.File, k, v); err != nil {
			return err
		}
	}
	return nil
}
