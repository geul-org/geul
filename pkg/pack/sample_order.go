//ff:func feature=pack type=helper control=sequence
//ff:what sampleOrder — the records in SHA-256 order: a fixed, content-spread pick for dictionary samples
package pack

import (
	"bytes"
	"slices"
)

// sampleOrder sorts a copy of ps by their SHA-256, so samples do not follow
// key order (which alternates version lists and documents).
func sampleOrder(ps []*put) []*put {
	out := slices.Clone(ps)
	slices.SortFunc(out, func(a, b *put) int { return bytes.Compare(a.sum[:], b.sum[:]) })
	return out
}
