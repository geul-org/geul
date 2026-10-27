//ff:type feature=notation type=model
//ff:what markerTokens — the slot marker tokens: ABSENT, UNKNOWN UNKNOWN-VALUE, NONE NO-VALUE
package notation

import "github.com/geul-org/gogeul/pkg/core"

// markerTokens maps the token of each named slot marker (reference/slot-marker)
// to its WORD. Reserved markers 0x0003~0x0007 have no token and are written
// as raw words.
var markerTokens = map[string]core.Word{"ABSENT": core.Absent, "UNKNOWN": core.UnknownValue, "NONE": core.NoValue}
