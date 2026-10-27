//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what TestScanStateStep — quotes, escapes, URIs and parentheses nest; brackets inside quotes and URIs do not count
package notation

import "testing"

func TestScanStateStep(t *testing.T) {
	s := &scanState{}
	for _, r := range `("a\"(<"<(>)` {
		s.step(r)
	}
	if !s.top() {
		t.Fatalf("%+v", s)
	}
}
