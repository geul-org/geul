//ff:func feature=schema type=rule control=sequence
//ff:what checkTarget — R2: a Reference target ID is present and in the fewest words
package schema

import "github.com/geul-org/geul/go/pkg/core"

// checkTarget applies R2 to the target ID words of a Reference.
func checkTarget(r rec, ws []core.Word) []Fact {
	if !minimalUnsigned(ws) {
		return []Fact{factAt("R2", Error, r, "대상 ID", "1 WORD 이상, 선행 0x0000 없음", fmtWords(len(ws)))}
	}
	return nil
}
