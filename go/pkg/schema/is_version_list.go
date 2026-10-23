//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what isVersionList — whether DOC-ID names document ID 0 (a version list document)
package schema

import "github.com/geul-org/geul/go/pkg/core"

// isVersionList reports whether the HEAD's DOC-ID has document ID 0, which is
// reserved for version list documents (edge/PUBLICATION).
func isVersionList(head []rec) bool {
	for _, r := range head {
		p := r.R.Payload
		if r.R.Codec == core.MetaDocID && len(p) >= 5 {
			return p[len(p)-4]|p[len(p)-3]|p[len(p)-2]|p[len(p)-1] == 0
		}
	}
	return false
}
