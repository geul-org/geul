//ff:func feature=pack type=helper control=sequence
//ff:what TestFileInfoMode — read-only bits; ModeDir for directories
package pack

import (
	"io/fs"
	"testing"
)

func TestFileInfoMode(t *testing.T) {
	if (fileInfo{dir: true}).Mode() != fs.ModeDir|0o555 || (fileInfo{}).Mode() != 0o444 {
		t.Fatal("mode")
	}
}
