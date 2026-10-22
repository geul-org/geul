//ff:func feature=inspect type=helper control=sequence
//ff:what TestAsData — the value when ok, nil otherwise
package inspect

import "testing"

func TestAsData(t *testing.T) {
	if got := asData(Widths{Edge: 2}, true); got != (Widths{Edge: 2}) {
		t.Fatalf("ok: %v", got)
	}
	if got := asData(Widths{Edge: 2}, false); got != nil {
		t.Fatalf("not ok: %v", got)
	}
}
