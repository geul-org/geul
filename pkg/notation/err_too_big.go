//ff:func feature=notation type=encoder control=sequence
//ff:what errTooBig — the error for an UNSIGNED-DECIMAL coefficient past int64
package notation

import "fmt"

// errTooBig reports s as too big for UNSIGNED-DECIMAL here.
func errTooBig(s string) error {
	return fmt.Errorf("geul/notation: UNSIGNED-DECIMAL %s: coefficient past 64 bits; write it as raw words", s)
}
