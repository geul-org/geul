//ff:func feature=pack type=writer control=sequence
//ff:what runReader.advance — read the next entry of the run into cur; ok false at the end
package pack

import (
	"errors"
	"fmt"
	"io"
)

// advance loads the next entry. A run that ends inside an entry is corrupt.
func (rr *runReader) advance() error {
	var b [hEntrySize]byte
	_, err := io.ReadFull(rr.r, b[:])
	switch {
	case errors.Is(err, io.EOF):
		rr.ok = false
		return nil
	case err != nil:
		return fmt.Errorf("%w: H run %s: %v", ErrCorrupt, rr.f.Name(), err)
	}
	rr.cur, rr.ok = parseHEntry(b[:]), true
	return nil
}
