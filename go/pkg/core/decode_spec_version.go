//ff:func feature=core type=decoder control=sequence
//ff:what DecodeSpecVersion — the spec version of a Meta SPEC-VERSION (minimal unsigned payload)
package core

import "fmt"

// DecodeSpecVersion returns the version r declares. r must be Meta SPEC-VERSION
// whose payload is an unsigned integer in minimal words that fits uint64.
func DecodeSpecVersion(r Record) (uint64, error) {
	if r.Kind != KindMeta || r.Codec != MetaSpecVersion {
		return 0, fmt.Errorf("geul/core: expected Meta SPEC-VERSION, got kind %d codec 0x%02X", r.Kind, r.Codec)
	}
	return wordsToUint(r.Payload)
}
