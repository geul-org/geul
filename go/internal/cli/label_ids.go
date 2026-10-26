//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what labelIDs — the Edge Id each #label of a batch stands for: its line's place after next
package cli

import (
	"strconv"
	"strings"
)

// labelIDs maps the leading `#n` of line i to next+i. Lines without a label
// map nothing.
func labelIDs(lines []string, next uint64) map[uint64]uint64 {
	out := map[uint64]uint64{}
	for i, l := range lines {
		first, _, _ := strings.Cut(l, " ")
		if n, err := strconv.ParseUint(strings.TrimPrefix(first, "#"), 10, 64); err == nil && strings.HasPrefix(first, "#") {
			out[n] = next + uint64(i)
		}
	}
	return out
}
