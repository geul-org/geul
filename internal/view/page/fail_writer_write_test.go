//ff:func feature=page type=helper control=sequence
//ff:what failWriter.Write — always fails
package page

import "errors"

// Write fails without writing.
func (failWriter) Write([]byte) (int, error) { return 0, errors.New("closed") }
