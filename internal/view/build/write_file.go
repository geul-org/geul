//ff:func feature=build type=helper control=sequence
//ff:what writeFile — create a file and fill it by a writer function through a buffer; the first error of writing, flushing or closing
package build

import (
	"bufio"
	"io"
	"os"
)

// writeFile creates (or truncates) the file at p and lets write fill it
// through a buffer (templates write in many small pieces).
func writeFile(p string, write func(io.Writer) error) error {
	f, err := os.Create(p)
	if err != nil {
		return err
	}
	bw := bufio.NewWriterSize(f, 64<<10)
	err = write(bw)
	if err == nil {
		err = bw.Flush()
	}
	if cerr := f.Close(); err == nil {
		err = cerr
	}
	return err
}
