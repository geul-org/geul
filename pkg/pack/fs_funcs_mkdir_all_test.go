//ff:func feature=pack type=helper control=sequence
//ff:what fsFuncs.MkdirAll — fail when the countdown reaches zero
package pack

import "errors"

// MkdirAll fails once mkdir counts down to zero.
func (f *fsFuncs) MkdirAll(dir string) error {
	if f.mkdir == 0 {
		return errors.New("mkdir failed")
	}
	f.mkdir--
	return osFS{}.MkdirAll(dir)
}
