//ff:func feature=schema type=decoder control=sequence
//ff:what edgeID — the Edge Id: the LOCAL Reference (Quantifier 000) leading an Edge payload
package schema

import "github.com/geul-org/geul/go/pkg/core"

// edgeID reads the LOCAL Reference that must lead an Edge payload (E1).
func edgeID(r rec) (uint64, bool) {
	ref, _, err := core.DecodeRecord(r.R.Payload)
	if err != nil {
		return 0, false
	}
	id, err := core.DecodeLocalRef(ref)
	return id, err == nil
}
