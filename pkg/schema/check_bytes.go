//ff:func feature=schema type=rule control=sequence
//ff:what checkBytes — V4 and V5: odd-codec padding, then the text rule of the codec pair
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// checkBytes applies V4 to a byte-string Value of an even/odd codec pair, then
// V5 to its bytes.
func checkBytes(r rec) []Fact {
	p := r.R.Payload
	b := core.PutWords(p)
	if r.R.Codec&1 == 1 && len(p) == 0 {
		return []Fact{factAt("V4", Error, r, "Payload", "홀수 코드는 1 WORD 이상", "0 WORD")}
	}
	if r.R.Codec&1 == 1 && b[len(b)-1] != 0 {
		return []Fact{factAt("V4", Error, r, "Payload", "마지막 하위 byte 0x00", fmt.Sprintf("0x%02X", b[len(b)-1]))}
	}
	if r.R.Codec&1 == 1 {
		b = b[:len(b)-1]
	}
	return checkText(r, r.R.Codec&^1, b)
}
