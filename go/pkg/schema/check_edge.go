//ff:func feature=schema type=rule control=sequence
//ff:what checkEdge — E1, E3 and E4 for one Edge record
package schema

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// checkEdge applies E1 (leading LOCAL Edge Id), E3 (reserved codec) and E4
// (layout of the core-embedded version list codecs).
func checkEdge(r rec) []Fact {
	var fs []Fact
	if _, ok := edgeID(r); !ok {
		fs = append(fs, factAt("E1", Error, r, "Edge Id", "Payload 첫 요소 LOCAL Reference (Quantifier 000)", "없음"))
	}
	if !edgeDefined(r.R.Codec) {
		fs = append(fs, factAt("E3", Warn, r, "Codec", "배정된 Edge 코드", fmt.Sprintf("0x%02X (예약)", r.R.Codec)))
	}
	if r.R.Codec >= core.EdgePublication && r.R.Codec <= core.EdgeTombstone {
		if _, err := core.DecodePublication(r.R); err != nil {
			fs = append(fs, factAt("E4", Error, r, "Payload", "edge/PUBLICATION 레이아웃", err.Error()))
		}
	}
	return fs
}
