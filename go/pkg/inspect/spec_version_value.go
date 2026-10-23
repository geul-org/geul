//ff:func feature=inspect type=decoder control=sequence
//ff:what specVersionValue — Meta SPEC-VERSION as an UNSIGNED-INTEGER Value
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// specVersionValue reads the bare unsigned integer of Meta SPEC-VERSION and
// presents it as an UNSIGNED-INTEGER Value. ok is false for a malformed payload.
func specVersionValue(r core.Record) (Value, bool) {
	v, err := core.DecodeSpecVersion(r)
	return Value{Codec: core.CodecUnsignedInteger, Uint: v}, err == nil
}
