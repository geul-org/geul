//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what removeStrayRuns — delete run files the checkpoint does not list (spilled after it)
package pack

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
)

// removeStrayRuns removes every run-* file in out's working folder that runs
// does not name: they hold H keys of records put after the checkpoint, which
// the caller puts again.
func removeStrayRuns(out string, runs []string) error {
	dir := filepath.Join(out, streamDir)
	ents, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, e := range ents {
		if !strings.HasPrefix(e.Name(), "run-") || slices.Contains(runs, e.Name()) {
			continue
		}
		if err := os.Remove(filepath.Join(dir, e.Name())); err != nil {
			return err
		}
	}
	return nil
}
