//ff:func feature=resolve type=rule control=sequence
//ff:what addPublication — append a PUBLICATION's version, which must follow the last one
package resolve

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// addPublication appends p's version to lg. Version IDs rise strictly.
func addPublication(lg *versionLog, p core.Publication) error {
	if n := len(lg.Entries); n > 0 && p.Version <= lg.Entries[n-1].Version {
		return fmt.Errorf("version %s out of order", FormatVersion(p.Version))
	}
	lg.Entries = append(lg.Entries, VersionEntry{Version: p.Version, Hash: p.Hash})
	return nil
}
