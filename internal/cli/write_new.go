//ff:func feature=cli type=helper control=sequence
//ff:what writeNew — write a file, refusing to replace an existing one unless forced
package cli

import (
	"fmt"
	"os"
)

// writeNew writes b to path; an existing path is kept unless force.
func writeNew(path string, b []byte, force bool) error {
	if _, err := os.Stat(path); err == nil && !force {
		return fmt.Errorf("%s exists (use --force to replace it)", path)
	}
	return os.WriteFile(path, b, 0o644)
}
