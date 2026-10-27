//ff:func feature=resolve type=rule control=iteration dimension=1
//ff:what markVersion — apply a DEPRECATION or TOMBSTONE to the version it names
package resolve

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// markVersion flags the version p names as deprecated or deleted. Naming a
// version not published earlier in the list is an error.
func markVersion(lg *versionLog, p core.Publication) error {
	for i := range lg.Entries {
		if lg.Entries[i].Version != p.Version {
			continue
		}
		lg.Entries[i].Deprecated = lg.Entries[i].Deprecated || p.Codec == core.EdgeDeprecation
		lg.Entries[i].Deleted = lg.Entries[i].Deleted || p.Codec == core.EdgeTombstone
		return nil
	}
	return fmt.Errorf("version %s not published", FormatVersion(p.Version))
}
