//ff:func feature=cli type=command control=sequence
//ff:what Execute — run geul with the process's arguments, environment and streams; the exit code
package cli

import (
	"fmt"
	"os"

	"github.com/geul-org/gogeul/internal/config"
)

// Execute runs the command line and returns 0, or 1 after printing the
// error to standard error.
func Execute() int {
	cwd, _ := os.Getwd()
	confDir, _ := os.UserConfigDir()
	cacheDir, _ := os.UserCacheDir()
	a := &app{env: config.Env{Getenv: os.Getenv, Cwd: cwd, ConfigDir: confDir, CacheDir: cacheDir}, in: os.Stdin, out: os.Stdout}
	if err := newRootCmd(a).Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "geul:", err)
		return 1
	}
	return 0
}
