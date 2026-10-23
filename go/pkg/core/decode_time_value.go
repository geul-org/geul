//ff:func feature=core type=decoder control=selection
//ff:what DecodeTimeValue — a GREGORIAN-DATE or POSIX-TIMESTAMP-SECONDS Value as a UTC instant
package core

import (
	"fmt"
	"time"
)

// DecodeTimeValue reads r as an instant (GEUL t002 docs/value/time): a
// GREGORIAN-DATE (2-word signed day count from 1970-01-01, read as that day's
// midnight UTC) or a POSIX-TIMESTAMP-SECONDS (4-word signed second count). Both
// are fixed width; any other codec or width is an error.
func DecodeTimeValue(r Record) (time.Time, error) {
	switch {
	case r.Kind == KindValue && r.Codec == CodecGregorianDate && len(r.Payload) == 2:
		days := int32(uint32(r.Payload[0])<<16 | uint32(r.Payload[1]))
		return time.Unix(int64(days)*86400, 0).UTC(), nil
	case r.Kind == KindValue && r.Codec == CodecPosixSeconds && len(r.Payload) == 4:
		s := int64(uint64(r.Payload[0])<<48 | uint64(r.Payload[1])<<32 | uint64(r.Payload[2])<<16 | uint64(r.Payload[3]))
		return time.Unix(s, 0).UTC(), nil
	default:
		return time.Time{}, fmt.Errorf("geul/core: not a GREGORIAN-DATE or POSIX-TIMESTAMP-SECONDS Value (kind %d codec 0x%02X, %d words)", r.Kind, r.Codec, len(r.Payload))
	}
}
