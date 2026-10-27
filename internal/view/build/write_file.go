//ff:func feature=build type=helper control=sequence
//ff:what writeFile — create a file and fill it by a writer function; the first error of writing or closing
package build

import (
	"io"
	"os"
)

// writeFile creates (or truncates) the file at p and lets write fill it.
func writeFile(p string, write func(io.Writer) error) error {
	f, err := os.Create(p)
	if err != nil {
		return err
	}
	err = write(f)
	if cerr := f.Close(); err == nil {
		err = cerr
	}
	return err
}
