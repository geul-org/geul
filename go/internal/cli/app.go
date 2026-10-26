//ff:type feature=cli type=model
//ff:what app — one geul run: where settings come from, the settings in effect, and the streams it reads and writes
package cli

import (
	"io"

	"github.com/geul-org/geul/go/internal/config"
)

// app is shared by every command of one run. env is read once; cfg is loaded
// before a command runs (load), after the flags are parsed.
type app struct {
	env config.Env
	cfg config.Config
	in  io.Reader
	out io.Writer
}
