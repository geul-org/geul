//ff:func feature=cli type=command control=iteration dimension=1
//ff:what app.check — collect findings of the given paths (and the archive), print them, fail when there are any
package cli

import (
	"errors"
	"fmt"
)

// check runs schema on every `.geul` of args (checkPath) and, with archive,
// on the data folder and its versions (checkArchive). Any finding fails.
func (a *app) check(args []string, archive bool) error {
	if len(args) == 0 && !archive {
		return errors.New("check: give files or folders, or --archive")
	}
	var fs []finding
	for _, p := range args {
		got, err := checkPath(p)
		if err != nil {
			return err
		}
		fs = append(fs, got...)
	}
	if archive {
		got, err := a.checkArchive()
		if err != nil {
			return err
		}
		fs = append(fs, got...)
	}
	if err := a.printFindings(fs); err != nil || len(fs) == 0 {
		return err
	}
	return fmt.Errorf("check: %d finding(s)", len(fs))
}
