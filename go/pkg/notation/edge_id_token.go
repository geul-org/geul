//ff:func feature=notation type=decoder control=sequence
//ff:what edgeIDToken — the #n token of an Edge's leading LOCAL Edge Id part; false when the first part is not one
package notation

import (
	"strings"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/inspect"
)

// edgeIDToken reads parts[0] as the Edge Id.
func edgeIDToken(parts []inspect.Part) (string, bool) {
	if len(parts) == 0 || !parts[0].Record {
		return "", false
	}
	r, _, _ := core.DecodeRecord(parts[0].Words)
	s := formatRecord(r)
	return s, strings.HasPrefix(s, "#")
}
