//ff:func feature=build type=helper control=sequence
//ff:what checkOut — refuse an output tree that is the source tree or inside it
package build

import (
	"fmt"
	"path/filepath"
	"strings"
)

// checkOut refuses to write the output tree out over the source tree src or
// inside it: the copies would overwrite the sources, or the walk would read
// its own output.
func checkOut(src, out string) error {
	s, _ := filepath.Abs(src)
	o, _ := filepath.Abs(out)
	if o == s || strings.HasPrefix(o, s+string(filepath.Separator)) {
		return fmt.Errorf("output %s is inside the source tree %s", out, src)
	}
	return nil
}
