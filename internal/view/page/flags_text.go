//ff:func feature=page type=helper control=iteration dimension=1
//ff:what flagsText — the set LINK flags by name (once, strike, ordered); gate and force are shown on their own
package page

import (
	"strings"

	"github.com/geul-org/gogeul/pkg/toulmin"
)

// linkFlagNames lists the LINK flags shown as words, in bit order.
var linkFlagNames = []struct {
	bit  toulmin.EdgeFlags
	name string
}{{toulmin.EdgeFlagOnce, "once"}, {toulmin.EdgeFlagStrike, "strike"}, {toulmin.EdgeFlagOrdered, "ordered"}}

// flagsText names the set once/strike/ordered flags, space-separated; "" when
// none is set. F_GATE and F_FORCE are not named: the gate and the attack force
// are their own fields.
func flagsText(flags toulmin.EdgeFlags) string {
	var names []string
	for _, fl := range linkFlagNames {
		if flags&fl.bit != 0 {
			names = append(names, fl.name)
		}
	}
	return strings.Join(names, " ")
}
