//ff:func feature=pack type=helper control=sequence
//ff:what badFile.WriteAt — fail when flagged
package pack

import "errors"

// WriteAt fails when f.writeAt is set.
func (f *badFile) WriteAt(b []byte, off int64) (int, error) {
	if f.writeAt {
		return 0, errors.New("writeAt failed")
	}
	return f.objFile.WriteAt(b, off)
}
