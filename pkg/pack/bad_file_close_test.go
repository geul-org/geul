//ff:func feature=pack type=helper control=sequence
//ff:what badFile.Close — close the real file, then fail when flagged
package pack

import "errors"

// Close fails when f.close is set (the real file is still closed).
func (f *badFile) Close() error {
	err := f.objFile.Close()
	if f.close {
		return errors.New("close failed")
	}
	return err
}
