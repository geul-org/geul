//ff:func feature=pack type=helper control=sequence
//ff:what TestDirFileRead — reading a directory fails
package pack

import (
	"errors"
	"io/fs"
	"testing"
)

func TestDirFileRead(t *testing.T) {
	if _, err := (&dirFile{}).Read(nil); !errors.Is(err, fs.ErrInvalid) {
		t.Fatal(err)
	}
}
