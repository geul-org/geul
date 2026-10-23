//ff:func feature=resolve type=encoder control=iteration dimension=1
//ff:what formatVersionList — entries as version list text, one line each
package resolve

import (
	"encoding/hex"
	"strings"
)

// formatVersionList writes list in the form parseVersionList reads.
func formatVersionList(list []VersionEntry) string {
	var b strings.Builder
	for _, e := range list {
		b.WriteString(FormatVersion(e.Version) + " 0x" + hex.EncodeToString(e.Hash[:]))
		if e.Deprecated {
			b.WriteString(" deprecated")
		}
		b.WriteString("\n")
	}
	return b.String()
}
