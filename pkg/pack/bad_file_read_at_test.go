//ff:func feature=pack type=helper control=sequence
//ff:what badFile.ReadAt — fail when flagged
package pack

import "errors"

// ReadAt fails when f.readAt is set.
func (f *badFile) ReadAt(b []byte, off int64) (int, error) {
	if f.readAt {
		return 0, errors.New("readAt failed")
	}
	return f.objFile.ReadAt(b, off)
}
