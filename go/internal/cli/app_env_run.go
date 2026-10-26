//ff:func feature=cli type=command control=selection
//ff:what app.envRun — geul env: list, look up, write or remove settings
package cli

import "errors"

// envRun writes (w: KEY=VALUE args) or removes (u: KEY args) settings in the
// settings file, or prints the named settings' values, or every setting with
// its source.
func (a *app) envRun(args []string, w, u bool) error {
	switch {
	case w && u:
		return errors.New("env: -w and -u together")
	case w || u:
		return a.envWrite(args, w)
	case len(args) > 0:
		return a.envValues(args)
	}
	return a.envList()
}
