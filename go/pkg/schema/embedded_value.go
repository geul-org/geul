//ff:func feature=schema type=decoder control=sequence
//ff:what embeddedValue — the single Value record a Meta payload carries
package schema

import "github.com/geul-org/geul/go/pkg/core"

// embeddedValue decodes m's payload as exactly one Value record.
func embeddedValue(m rec) (rec, bool) {
	v, n, err := core.DecodeRecord(m.R.Payload)
	if err != nil || n != len(m.R.Payload) || v.Kind != core.KindValue {
		return rec{}, false
	}
	return rec{Index: m.Index, Word: m.Word, Total: n, R: v}, true
}
