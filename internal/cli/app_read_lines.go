//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what app.readLines — the lines of a list file or standard input, without blank lines and // comments
package cli

import (
	"io"
	"os"
	"strings"
)

// readLines reads path (`-` is standard input) and keeps the lines that are
// not blank and do not start with `//`.
func (a *app) readLines(path string) ([]string, error) {
	var b []byte
	var err error
	if path == "-" {
		b, err = io.ReadAll(a.in)
	} else {
		b, err = os.ReadFile(path)
	}
	if err != nil {
		return nil, err
	}
	var out []string
	for _, l := range strings.Split(string(b), "\n") {
		if l = strings.TrimSpace(l); l != "" && !strings.HasPrefix(l, "//") {
			out = append(out, l)
		}
	}
	return out, nil
}
