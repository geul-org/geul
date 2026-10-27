//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what Dir.removeOtherVersions — deletes a document's version files other than version 1
package resolve

import (
	"os"
	"path/filepath"
	"strings"
)

// removeOtherVersions deletes every `0x????????.geul` version file in the
// folder of document k except version 1 (Overwrite). The version list and
// other files stay.
func (d Dir) removeOtherVersions(k DocKey) error {
	dir := d.file(k.Issuer + "/" + k.Doc)
	files, err := filepath.Glob(filepath.Join(dir, "0x????????.geul"))
	if err != nil {
		return err
	}
	keep := FormatVersion(1) + ".geul"
	for _, f := range files {
		if strings.EqualFold(filepath.Base(f), keep) {
			continue
		}
		if err := os.Remove(f); err != nil {
			return err
		}
	}
	return nil
}
