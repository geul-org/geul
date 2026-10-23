//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what parseVersionList — a version list text as entries in strictly ascending version order
package resolve

import (
	"fmt"
	"strings"
)

// parseVersionList reads a version list: one parseVersionLine per non-empty
// line, versions strictly ascending. The format is provisional until GEUL-ORG
// fixes the signed list format.
func parseVersionList(b []byte) ([]VersionEntry, error) {
	var out []VersionEntry
	for _, line := range strings.Split(string(b), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		e, err := parseVersionLine(line)
		if err != nil {
			return nil, err
		}
		if len(out) > 0 && e.Version <= out[len(out)-1].Version {
			return nil, fmt.Errorf("geul/resolve: version %s out of order", FormatVersion(e.Version))
		}
		out = append(out, e)
	}
	return out, nil
}
