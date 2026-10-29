//ff:func feature=cli type=command control=sequence
//ff:what packUnpack — write a pack archive back as a directory tree under another folder (inspection, older tools)
package cli

import (
	"fmt"
	"io"
	"path/filepath"
)

// packUnpack restores dir's pack archive as out/.well-known/geul/.
func packUnpack(dir, out string, w io.Writer) error {
	a, err := openArchive(dir, packOptions())
	if err != nil {
		return err
	}
	n, err := a.Unpack(out)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(w, "geul: %d files → %s\n", n, filepath.Join(out, ".well-known", "geul"))
	return err
}
