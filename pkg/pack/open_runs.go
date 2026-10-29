//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what openRuns — open every spilled run of a stream positioned at its first entry
package pack

import (
	"bufio"
	"os"
	"path/filepath"
)

// openRuns opens the named runs in out's working folder, each at its first
// entry. On error the files opened so far are closed.
func openRuns(out string, names []string) ([]*runReader, error) {
	var rs []*runReader
	for _, n := range names {
		f, err := os.Open(filepath.Join(out, streamDir, n))
		if err == nil {
			rr := &runReader{f: f, r: bufio.NewReaderSize(f, 1<<20)}
			rs = append(rs, rr)
			err = rr.advance()
		}
		if err != nil {
			closeRuns(rs)
			return nil, err
		}
	}
	return rs, nil
}
