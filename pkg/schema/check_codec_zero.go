//ff:func feature=schema type=rule control=selection
//ff:what checkCodecZero — R1: a HEAD or BODY record of Reference Codec 0000 — a named slot marker is invalid there, a reserved marker or a Codec 0000 Reference with payload is reserved
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// checkCodecZero applies R1 to a record of the reserved Reference Codec 0000
// standing as a HEAD or BODY record. Its low 3 bits are a slot marker number,
// not a Quantifier. A named slot marker (ABSENT, UNKNOWN-VALUE, NO-VALUE) is
// valid only inside a record field that names it (reference/slot-marker), so
// here it is an Error; a reserved marker (0x0003~0x0007) or a Codec 0000
// Reference with a payload is reserved (Warn).
func checkCodecZero(r rec) []Fact {
	name := core.SlotMarkerName(r.R.Codec)
	switch {
	case core.IsSlotMarker(r.R) && name != "":
		return []Fact{factAt("R1", Error, r, "자리 표지", "표지를 받는 레코드 필드 안", fmt.Sprintf("레코드 자리의 %s (0x%04X)", name, r.R.Codec))}
	case core.IsSlotMarker(r.R):
		return []Fact{factAt("R1", Warn, r, "자리 표지", "0x0000~0x0002", fmt.Sprintf("0x%04X (예약 표지)", r.R.Codec))}
	default:
		return []Fact{factAt("R1", Warn, r, "Codec", "0001~1111", "0000 (예약)")}
	}
}
