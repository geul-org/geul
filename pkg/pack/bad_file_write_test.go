//ff:func feature=pack type=helper control=sequence
//ff:what badFile.Write — fail when flagged
package pack

import "errors"

// Write fails when f.write is set.
func (f *badFile) Write(b []byte) (int, error) {
	if f.write {
		return 0, errors.New("write failed")
	}
	return f.objFile.Write(b)
}
