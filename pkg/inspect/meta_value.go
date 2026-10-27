//ff:func feature=inspect type=decoder control=sequence
//ff:what metaValue — the Value of a Meta whose payload is exactly one known Value record
package inspect

import "github.com/geul-org/gogeul/pkg/core"

// metaValue decodes a Meta record whose payload is exactly one embedded Value
// of a known codec. ok is false when the payload is anything else.
func metaValue(r core.Record) (Value, bool) {
	v, n, err := core.DecodeRecord(r.Payload)
	if err != nil || n != len(r.Payload) || v.Kind != core.KindValue {
		return Value{}, false
	}
	return ParseValue(v)
}
