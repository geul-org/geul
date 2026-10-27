//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what Dir.basedOn — a live stored version of a document that is a diff on version v (DELTA-BASE)
package resolve

import "context"

// basedOn returns the first version of k, other than v and not deleted, whose
// stored file is a diff with DELTA-BASE v.
func (d Dir) basedOn(k DocKey, v uint32) (uint32, bool) {
	lg, err := d.readList(k, k.Doc)
	if err != nil {
		return 0, false
	}
	for _, e := range lg.Entries {
		if e.Version == v || e.Deleted {
			continue
		}
		b, err := d.Fetch(context.Background(), versionPath(k, e.Version))
		if base, ok := deltaBaseOf(b); err == nil && ok && base == v {
			return e.Version, true
		}
	}
	return 0, false
}
