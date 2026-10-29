//ff:func feature=pack type=encoder control=sequence
//ff:what failCodec.Compress — always an error
package pack

import "errors"

// Compress fails.
func (failCodec) Compress([]byte, *Dict) ([]byte, error) {
	return nil, errors.New("compress failed")
}
