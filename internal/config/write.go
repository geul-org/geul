//ff:func feature=config type=helper control=iteration dimension=1
//ff:what Write — set or remove one key in the settings file, keeping its other lines
package config

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

// Write sets key to value in the settings file at path (`geul env -w`), or
// removes it when value is empty (`geul env -u`). Other lines stay as they
// are. Only keys of Keys other than GEULENV can be written.
func Write(path, key, value string) error {
	if !slices.Contains(Keys, key) || key == "GEULENV" {
		return fmt.Errorf("%s is not a setting the file can hold", key)
	}
	b, err := os.ReadFile(path)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	var out []string
	for _, line := range strings.Split(strings.TrimRight(string(b), "\n"), "\n") {
		k, _, _ := strings.Cut(line, "=")
		if line != "" && strings.TrimSpace(k) != key {
			out = append(out, line)
		}
	}
	if value != "" {
		out = append(out, key+"="+value)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(strings.Join(out, "\n")+"\n"), 0o644)
}
