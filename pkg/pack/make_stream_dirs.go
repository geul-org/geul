//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what makeStreamDirs — create a set folder's o/, m/ and the stream's working folder
package pack

import (
	"os"
	"path/filepath"
)

// makeStreamDirs creates out/o, out/m and out/.stream.
func makeStreamDirs(out string) error {
	for _, d := range []string{"o", "m", streamDir} {
		if err := os.MkdirAll(filepath.Join(out, d), 0o755); err != nil {
			return err
		}
	}
	return nil
}
