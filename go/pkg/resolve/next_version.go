//ff:func feature=resolve type=helper control=sequence
//ff:what nextVersion — the version ID after the list's last one; false when 2 words are used up
package resolve

import "math"

// nextVersion is the version ID a new publication gets: 1 for an empty list,
// otherwise one past the last. When the last is the 2-word maximum the document
// is exhausted and must be published under a new document ID (GEUL-ORG 버전 ID).
func nextVersion(list []VersionEntry) (uint32, bool) {
	if len(list) == 0 {
		return 1, true
	}
	last := list[len(list)-1].Version
	return last + 1, last != math.MaxUint32
}
