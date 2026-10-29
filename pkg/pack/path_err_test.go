//ff:func feature=pack type=helper control=sequence
//ff:what TestPathErr — an *fs.PathError wrapping the cause
package pack

import (
	"errors"
	"io/fs"
	"testing"
)

func TestPathErr(t *testing.T) {
	err := pathErr("read", "a", fs.ErrNotExist)
	var pe *fs.PathError
	if !errors.As(err, &pe) || pe.Op != "read" || pe.Path != "a" || !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
}
