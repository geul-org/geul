//ff:type feature=pack type=model
//ff:what runReader — one H run file read entry by entry during the merge
package pack

import (
	"bufio"
	"os"
)

// runReader reads one sorted run. cur is the entry at the front; ok is false
// once the run is used up.
type runReader struct {
	f   *os.File
	r   *bufio.Reader
	cur hEntry
	ok  bool
}
