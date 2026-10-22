//ff:func feature=geulmd type=encoder control=iteration dimension=1
//ff:what tidList — Edge Ids as `#a,#b` (empty list as `#`)
package geulmd

import (
	"strconv"
	"strings"

	"github.com/geul-org/geul/go/pkg/toulmin"
)

// tidList renders a REF list (F_WITH, F_GROUNDS) as comma-joined `#id`s; an
// empty but present list is a lone `#`, so presence survives.
func tidList(ids []toulmin.TID) string {
	parts := make([]string, 0, len(ids))
	for _, id := range ids {
		parts = append(parts, strconv.FormatUint(uint64(id), 10))
	}
	return "#" + strings.Join(parts, ",#")
}
