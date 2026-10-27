//ff:func feature=cli type=helper control=sequence
//ff:what failWriter.Write — always fails
package cli

import "errors"

// Write fails.
func (failWriter) Write([]byte) (int, error) { return 0, errors.New("closed") }
