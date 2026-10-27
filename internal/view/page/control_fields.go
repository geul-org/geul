//ff:func feature=page type=builder control=selection
//ff:what controlFields — a SECTION's section name or a CASE-BEGIN/END's case link
package page

import (
	"strconv"

	"github.com/geul-org/gogeul/pkg/toulmin"
)

// sectionNames names the TANGEUL section codes (GEUL t002 TOULMIN 섹션 코드).
var sectionNames = map[uint64]string{1: "Subject", 2: "See", 3: "Definitions", 4: "Rules", 5: "Cases", 6: "Provides", 7: "Internal"}

// controlFields lays out a framing record: SECTION its section (by name when
// known), CASE-BEGIN/CASE-END a link to its case. PROSE has only its span.
func controlFields(p toulmin.ControlPacket, labels map[uint64]string) []Field {
	switch p.Code {
	case toulmin.ControlSection:
		name, ok := sectionNames[p.Arg]
		if !ok {
			name = strconv.FormatUint(p.Arg, 10)
		}
		return []Field{{Key: "section", Text: name}}
	case toulmin.ControlCaseBegin, toulmin.ControlCaseEnd:
		return []Field{{Key: "case", Refs: refs([]uint64{p.Arg}, labels)}}
	default:
		return nil
	}
}
