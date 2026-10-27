//ff:func feature=schema type=rule control=sequence
//ff:what checkSpecVersion — M4: SPEC-VERSION's payload is a bare minimal unsigned integer naming the spec this checker knows
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// checkSpecVersion applies M4 to Meta SPEC-VERSION: the payload is the version
// itself, a minimal unsigned integer with no embedded Value (t002 is 0x0002).
// Another version is a Warn: the document may follow a spec this checker
// predates.
func checkSpecVersion(r rec) []Fact {
	v, err := core.DecodeSpecVersion(r.R)
	if err != nil {
		return []Fact{factAt("M4", Error, r, "Payload", "가장 짧은 폭의 unsigned 정수 (t002 = 0x0002)", fmtWords(len(r.R.Payload)))}
	}
	if v != core.SpecVersion {
		return []Fact{factAt("M4", Warn, r, "Payload", fmt.Sprintf("0x%04x", core.SpecVersion), fmt.Sprintf("0x%04x (모르는 명세 판)", v))}
	}
	return nil
}
