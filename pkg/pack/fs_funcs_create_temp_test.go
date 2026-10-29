//ff:func feature=pack type=helper control=sequence
//ff:what fsFuncs.CreateTemp — fail when the countdown reaches zero, else a real temp file (wrapped when asked)
package pack

import "errors"

// CreateTemp fails once create counts down to zero.
func (f *fsFuncs) CreateTemp(dir string) (objFile, error) {
	if f.create == 0 {
		return nil, errors.New("create failed")
	}
	f.create--
	o, err := osFS{}.CreateTemp(dir)
	if err == nil && f.wrap != nil {
		o = f.wrap(o)
	}
	return o, err
}
