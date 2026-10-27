//ff:type feature=resolve type=model
//ff:what timeValue — an instant read from a time Meta
package resolve

import "time"

// timeValue is an instant read from EFFECTIVE-FROM or EFFECTIVE-UNTIL.
type timeValue struct {
	time.Time
}
