//ff:func feature=delta type=encoder control=sequence
//ff:what TestFresh — consecutive new IDs from the auto increment, marked used
package delta

import "testing"

func TestFresh(t *testing.T) {
	l, _ := NewLocalIDs(nil, 7)
	if a, b := l.fresh(), l.fresh(); a != 7 || b != 8 || !l.used[7] || l.Next() != 9 {
		t.Fatalf("%d %d", a, b)
	}
}
