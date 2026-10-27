//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what codecDecls — E5: each CODEC names a FIELD of the document, a target once per field, a raw field has at most one exact and no Absent
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// codecDecls checks the CODEC Edges of body against its FIELDs (edge/CODEC):
// the field is a FIELD Edge of this document, no field names the same target
// twice, a raw field (width ≥ 1) has at most one CODEC, an exact one other
// than the Absent target 0000 (a raw field has no TAG to leave out), and no
// field has 0000 alone (absentOnly). CODECs that do not decode are E4's.
func codecDecls(body []rec) []Fact {
	widths := map[uint64]core.Word{}
	for _, r := range body {
		if f, err := core.DecodeFieldDecl(r.R); err == nil {
			widths[f.EdgeID] = f.Width
		}
	}
	seen := map[core.CodecDecl]bool{}
	count := map[uint64]int{}
	var fs []Fact
	for _, r := range body {
		c, err := core.DecodeCodecDecl(r.R)
		if err != nil {
			continue
		}
		w, ok := widths[c.Field]
		key := core.CodecDecl{Field: c.Field, Target: c.Target}
		count[c.Field]++
		switch {
		case !ok:
			fs = append(fs, factAt("E5", Error, r, "필드", "같은 문서의 FIELD", fmt.Sprintf("#%d", c.Field)))
		case seen[key]:
			fs = append(fs, factAt("E5", Error, r, "대상", "필드마다 한 번", fmt.Sprintf("%04X 중복", c.Target)))
		case w != 0 && c.Target == core.Absent:
			fs = append(fs, factAt("E5", Error, r, "대상", "날값 필드는 0000(없음) 불가", "0000"))
		case w != 0 && c.Target&0x0780 != 0:
			fs = append(fs, factAt("E5", Error, r, "대상", "날값 필드는 정확한 대상", fmt.Sprintf("%04X 와일드카드", c.Target)))
		case w != 0 && count[c.Field] > 1:
			fs = append(fs, factAt("E5", Error, r, "필드", "날값 필드의 CODEC 하나", fmt.Sprintf("#%d에 %d개", c.Field, count[c.Field])))
		}
		seen[key] = true
	}
	return append(fs, absentOnly(body)...)
}
