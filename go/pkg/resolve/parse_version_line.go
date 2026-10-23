//ff:func feature=resolve type=decoder control=sequence
//ff:what parseVersionLine — `<버전 ID> <내용 해시> [deprecated]` as a VersionEntry
package resolve

import (
	"encoding/binary"
	"fmt"
	"strings"
)

// parseVersionLine reads one version list line: the 2-word version ID (not 0),
// the 16-word content hash and an optional `deprecated` mark.
func parseVersionLine(s string) (VersionEntry, error) {
	f := strings.Fields(s)
	if len(f) < 2 || len(f) > 3 || (len(f) == 3 && f[2] != "deprecated") {
		return VersionEntry{}, fmt.Errorf("geul/resolve: bad version line %q", s)
	}
	v, ok := parseHexID(f[0], 2)
	h, okh := parseHexID(f[1], 16)
	if !ok || !okh || binary.BigEndian.Uint32(v) == 0 {
		return VersionEntry{}, fmt.Errorf("geul/resolve: bad version line %q", s)
	}
	e := VersionEntry{Version: binary.BigEndian.Uint32(v), Deprecated: len(f) == 3}
	copy(e.Hash[:], h)
	return e, nil
}
