//ff:func feature=pack type=encoder control=sequence
//ff:what failAfter.Compress — a raw frame while calls are left, then an error
package pack

import "errors"

// Compress frames src raw (Stored) until the calls run out.
func (f failAfter) Compress(src []byte, d *Dict) ([]byte, error) {
	if *f.left <= 0 {
		return nil, errors.New("compress failed")
	}
	*f.left--
	return Stored{}.Compress(src, d)
}
