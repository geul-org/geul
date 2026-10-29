//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamCountPut — Count is the number of P records put
package pack

import "testing"

func TestStreamCountPut(t *testing.T) {
	s := streamTestNew(t, nil)
	streamTestPut(t, s, 1, 2)
	if s.Count() != 4 {
		t.Fatal(s.Count())
	}
}
