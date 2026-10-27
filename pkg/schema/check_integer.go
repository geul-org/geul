//ff:func feature=schema type=rule control=sequence
//ff:what checkInteger — V3: SIGNED-/UNSIGNED-INTEGER in the fewest words
package schema

import "github.com/geul-org/gogeul/pkg/core"

// checkInteger applies V3 to an integer payload.
func checkInteger(r rec, ws []core.Word, signed bool, part string) []Fact {
	ok := minimalUnsigned(ws)
	if signed {
		ok = minimalSigned(ws)
	}
	if !ok {
		return []Fact{factAt("V3", Error, r, part, "1 WORD 이상의 최소 WORD 정수", fmtWords(len(ws)))}
	}
	return nil
}
