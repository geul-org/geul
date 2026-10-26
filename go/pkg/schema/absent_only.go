//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what absentOnly — E5: a field whose only CODEC is the Absent target 0000 (always left out)
package schema

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// absentOnly reports each field whose CODEC list is the Absent target 0000
// alone: such a field could never hold anything (edge/CODEC 선택 필드).
func absentOnly(body []rec) []Fact {
	targets := map[uint64]int{}
	absent := map[uint64]rec{}
	var order []uint64
	for _, r := range body {
		c, err := core.DecodeCodecDecl(r.R)
		if err != nil {
			continue
		}
		if targets[c.Field] == 0 {
			order = append(order, c.Field)
		}
		targets[c.Field]++
		if c.Target == core.Absent {
			absent[c.Field] = r
		}
	}
	var fs []Fact
	for _, f := range order {
		if r, ok := absent[f]; ok && targets[f] == 1 {
			fs = append(fs, factAt("E5", Error, r, "대상", "0000 말고 대상 하나 이상", fmt.Sprintf("#%d에 0000만", f)))
		}
	}
	return fs
}
