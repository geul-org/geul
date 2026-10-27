//ff:func feature=cli type=helper control=sequence
//ff:what parseID — an Edge Id written as #n
package cli

import (
	"fmt"
	"strconv"
	"strings"
)

// parseID reads `#n` (n > 0).
func parseID(s string) (uint64, error) {
	n, err := strconv.ParseUint(strings.TrimPrefix(s, "#"), 10, 64)
	if err != nil || !strings.HasPrefix(s, "#") || n == 0 {
		return 0, fmt.Errorf("%q is not an Edge Id #n", s)
	}
	return n, nil
}
