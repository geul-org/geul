//ff:type feature=cli type=model
//ff:what app — one geul run: where settings come from, the settings in effect, and the streams it reads and writes
package cli

import (
	"io"

	"github.com/geul-org/gogeul/internal/config"
	"github.com/geul-org/gogeul/pkg/pack"
)

// app is shared by every command of one run. env is read once; cfg is loaded
// before a command runs (load), after the flags are parsed. arch is the data
// folder's pack archive once opened (nil for a directory archive); its staged
// writes are committed after the command succeeds (commit).
type app struct {
	env  config.Env
	cfg  config.Config
	in   io.Reader
	out  io.Writer
	arch *pack.Archive
	dirs bool // the data folder was found to be a directory archive
}
