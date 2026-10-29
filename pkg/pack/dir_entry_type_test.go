//ff:func feature=pack type=helper control=sequence
//ff:what TestDirEntryType — ModeDir for directories, 0 for files
package pack

import (
	"io/fs"
	"testing"
)

func TestDirEntryType(t *testing.T) {
	if (dirEntry{dir: true}).Type() != fs.ModeDir || (dirEntry{}).Type() != 0 {
		t.Fatal("type")
	}
}
