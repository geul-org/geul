//ff:func feature=core type=decoder control=selection
//ff:what DecodeNextLocalID — the next LOCAL ID of a NEXT-LOCAL-ID Meta; an error for another record, a non-minimal payload or 0
package core

import (
	"errors"
	"fmt"
)

// DecodeNextLocalID reads a NEXT-LOCAL-ID Meta. The payload is a minimal
// unsigned integer of at least 1 (LOCAL ID 0 is never assigned).
func DecodeNextLocalID(r Record) (uint64, error) {
	if r.Kind != KindMeta || r.Codec != MetaNextLocalID {
		return 0, fmt.Errorf("geul/core: expected Meta NEXT-LOCAL-ID, got kind %d codec 0x%02X", r.Kind, r.Codec)
	}
	n, err := wordsToUint(r.Payload)
	switch {
	case err != nil:
		return 0, err
	case n == 0:
		return 0, errors.New("geul/core: NEXT-LOCAL-ID is 0")
	default:
		return n, nil
	}
}
