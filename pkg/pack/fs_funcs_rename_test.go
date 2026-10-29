//ff:func feature=pack type=helper control=sequence
//ff:what fsFuncs.Rename — fail when the countdown reaches zero
package pack

import "errors"

// Rename fails once rename counts down to zero.
func (f *fsFuncs) Rename(from, to string) error {
	if f.rename == 0 {
		return errors.New("rename failed")
	}
	f.rename--
	return osFS{}.Rename(from, to)
}
