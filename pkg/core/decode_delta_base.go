//ff:func feature=core type=decoder control=selection
//ff:what DecodeDeltaBase — the base version ID of a DELTA-BASE Meta; an error for another record, a payload other than 2 words or version 0
package core

import (
	"errors"
	"fmt"
)

// DecodeDeltaBase reads a DELTA-BASE Meta: exactly 2 words, a version ID of
// at least 1.
func DecodeDeltaBase(r Record) (uint32, error) {
	if r.Kind != KindMeta || r.Codec != MetaDeltaBase {
		return 0, fmt.Errorf("geul/core: expected Meta DELTA-BASE, got kind %d codec 0x%02X", r.Kind, r.Codec)
	}
	switch {
	case len(r.Payload) != 2:
		return 0, fmt.Errorf("geul/core: DELTA-BASE needs 2 words, got %d", len(r.Payload))
	case r.Payload[0] == 0 && r.Payload[1] == 0:
		return 0, errors.New("geul/core: DELTA-BASE version is 0")
	default:
		return uint32(r.Payload[0])<<16 | uint32(r.Payload[1]), nil
	}
}
