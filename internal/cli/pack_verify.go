//ff:func feature=cli type=command control=iteration dimension=1
//ff:what packVerify — verify every set of a pack archive and list the sets
package cli

import (
	"context"
	"fmt"
	"io"
)

// packVerify checks object hashes and index ↔ inventory agreement, then
// prints one line per set.
func packVerify(dir string, w io.Writer) error {
	a, err := openArchive(dir)
	if err != nil {
		return err
	}
	if err := a.Verify(context.Background()); err != nil {
		return err
	}
	for _, s := range a.Sets() {
		name := s.Name
		if name == "" {
			name = "."
		}
		fmt.Fprintf(w, "%s generation %d manifest %s objects %d ok\n", name, s.Generation, s.Manifest, s.Objects)
	}
	return nil
}
