//ff:func feature=resolve type=decoder control=sequence
//ff:what codecTargetOf — the target word a GEUL Issuer document ID stands for when it is a codec spec document ID
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// codecTargetOf reads a GEUL Issuer document ID in canonical hex and returns
// its target word when it is a codec spec document ID (codecDocID): three
// 0x0000 words, then a Kind·Codec word with LLLL = 0 that is not 0.
func codecTargetOf(docHex string) (core.Word, bool) {
	b, ok := parseHexID(docHex, 4)
	if !ok {
		return 0, false
	}
	ws, _ := core.ReadWords(b)
	t := ws[3]
	if ws[0]|ws[1]|ws[2] != 0 || t == 0 || t&0x0780 != 0 {
		return 0, false
	}
	return t, true
}
