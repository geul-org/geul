//ff:type feature=resolve type=model
//ff:what listTime — one version list of a PARENT chain: how many BODY records it had and when it was recorded
package resolve

import "time"

// listTime is one list document of a version list chain: n BODY records,
// recorded at at.
type listTime struct {
	n  int
	at time.Time
}
