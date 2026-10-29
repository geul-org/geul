//ff:func feature=cli type=helper control=sequence
//ff:what app.commit — after a command succeeds, write the pack archive's staged changes as new generations
package cli

import "context"

// commit records what the command wrote; a directory archive was written in place.
func (a *app) commit() error {
	if a.arch == nil {
		return nil
	}
	_, err := a.arch.Commit(context.Background())
	return err
}
