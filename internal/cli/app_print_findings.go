//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what app.printFindings — findings as `path: message` lines, or as JSON
package cli

import "fmt"

// printFindings writes fs.
func (a *app) printFindings(fs []finding) error {
	if done, err := a.json(fs); done {
		return err
	}
	for _, f := range fs {
		fmt.Fprintf(a.out, "%s: %s\n", f.Path, f.Message)
	}
	return nil
}
