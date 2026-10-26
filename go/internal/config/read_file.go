//ff:func feature=config type=helper control=iteration dimension=1
//ff:what readFile — the KEY=VALUE lines of a settings file; a missing file is empty
package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// readFile reads path: one `KEY=VALUE` per line, blank lines and `#`
// comments skipped (the format of `go env -w`). A missing file is no
// settings; a line without `=` is an error.
func readFile(path string) (map[string]string, error) {
	b, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return map[string]string{}, nil
	}
	if err != nil {
		return nil, err
	}
	out := map[string]string{}
	for i, line := range strings.Split(string(b), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		k, v, ok := strings.Cut(line, "=")
		if !ok {
			return nil, fmt.Errorf("geul: %s:%d: not KEY=VALUE", path, i+1)
		}
		out[strings.TrimSpace(k)] = strings.TrimSpace(v)
	}
	return out, nil
}
