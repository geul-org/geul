//ff:func feature=geulmd type=encoder control=iteration dimension=1
//ff:what hexWords — words as space-separated 4-digit lowercase hex
package geulmd

import (
	"fmt"
	"strings"

	"github.com/geul-org/geul/go/pkg/core"
)

// hexWords renders ws as space-separated `0x` words of exactly four lowercase
// hex digits, the raw form every record can fall back to.
func hexWords(ws []core.Word) string {
	parts := make([]string, 0, len(ws))
	for _, w := range ws {
		parts = append(parts, fmt.Sprintf("0x%04x", uint16(w)))
	}
	return strings.Join(parts, " ")
}
