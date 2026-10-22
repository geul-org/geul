//ff:func feature=core type=decoder control=sequence
//ff:what DecodeUintValue — the value of an UNSIGNED-INTEGER Value record
package core

import "fmt"

// DecodeUintValue returns the value of r, which must be an UNSIGNED-INTEGER
// Value record in minimal words that fits uint64.
func DecodeUintValue(r Record) (uint64, error) {
	if r.Kind != KindValue || r.Codec != CodecUnsignedInteger {
		return 0, fmt.Errorf("geul/core: expected UNSIGNED-INTEGER Value, got kind %d codec 0x%02X", r.Kind, r.Codec)
	}
	return wordsToUint(r.Payload)
}
