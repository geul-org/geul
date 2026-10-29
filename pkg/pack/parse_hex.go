//ff:func feature=pack type=decoder control=sequence
//ff:what parseHex — the bytes of `0x` + exactly digits lowercase hex digits, false otherwise
package pack

import (
	"encoding/hex"
	"strings"
)

// parseHex decodes canonical GEUL hex: `0x` then exactly digits lowercase digits.
func parseHex(s string, digits int) ([]byte, bool) {
	if len(s) != 2+digits || !strings.HasPrefix(s, "0x") || strings.ToLower(s) != s {
		return nil, false
	}
	b, err := hex.DecodeString(s[2:])
	return b, err == nil
}
