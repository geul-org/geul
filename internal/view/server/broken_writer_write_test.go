//ff:func feature=server type=helper control=sequence
//ff:what brokenWriter.Write — always fails
package server

import "errors"

// Write fails without writing.
func (brokenWriter) Write([]byte) (int, error) { return 0, errors.New("closed") }
