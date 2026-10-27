//ff:func feature=resolve type=decoder control=sequence
//ff:what parseHexID — bytes of a canonical `0x` lowercase hex ID of the given word width
package resolve

import (
	"encoding/hex"
	"strings"
)

// parseHexID decodes s, which must be `0x` followed by exactly words×4
// lowercase hex digits (GEUL-ORG ID 표기), and reports whether it was.
func parseHexID(s string, words int) ([]byte, bool) {
	if len(s) != 2+words*4 || !strings.HasPrefix(s, "0x") || strings.ToLower(s) != s {
		return nil, false
	}
	b, err := hex.DecodeString(s[2:])
	return b, err == nil
}
