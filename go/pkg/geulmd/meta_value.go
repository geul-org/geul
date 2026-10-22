//ff:func feature=geulmd type=encoder control=sequence
//ff:what metaValue — `NAME <literal>` for a Meta whose payload is exactly one known Value
package geulmd

import "github.com/geul-org/geul/go/pkg/core"

// metaValue renders a Meta record whose payload is exactly one embedded Value
// GEULMD knows as `name <literal>`. ok is false when the payload is anything
// else, so the caller falls back to a raw line.
func metaValue(name string, r core.Record) (string, bool) {
	v, n, err := core.DecodeRecord(r.Payload)
	if err != nil || n != len(r.Payload) {
		return "", false
	}
	lit, ok := renderValue(v)
	return name + " " + lit, ok
}
