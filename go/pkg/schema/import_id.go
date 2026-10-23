//ff:func feature=schema type=decoder control=iteration dimension=1
//ff:what importID — the IMPORT ID of an IMPORT-DOC-*, when it is well formed and fits 64 bits
package schema

import "github.com/geul-org/geul/go/pkg/core"

// importID reads the IMPORT ID of an IMPORT-DOC-* record.
func importID(r rec) (uint64, bool) {
	n := len(r.R.Payload) - core.ImportIssuerWidth(r.R.Codec) - 6
	if n < 1 || n > 4 {
		return 0, false
	}
	var v uint64
	for _, w := range r.R.Payload[:n] {
		v = v<<16 | uint64(w)
	}
	return v, true
}
