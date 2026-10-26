//ff:func feature=notation type=encoder control=sequence
//ff:what parseHex — raw words from 0x and a multiple of four hex digits
package notation

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/geul-org/geul/go/pkg/core"
)

// parseHex reads `0x` followed by 4 hex digits per word (either case).
func parseHex(tok string) ([]core.Word, error) {
	digits, ok := strings.CutPrefix(tok, "0x")
	b, err := hex.DecodeString(digits)
	if !ok || err != nil || len(digits) == 0 || len(digits)%4 != 0 {
		return nil, fmt.Errorf("geul/notation: %q is not 0x and 4 hex digits per word", tok)
	}
	ws, _ := core.ReadWords(b)
	return ws, nil
}
