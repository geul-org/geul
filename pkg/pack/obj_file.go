//ff:type feature=pack type=model
//ff:what objFile — the temp file an object is written to: append, patch the header, read back to hash
package pack

import "io"

// objFile is what *os.File gives the writer.
type objFile interface {
	io.Writer
	io.WriterAt
	io.ReaderAt
	Close() error
	Name() string
}
