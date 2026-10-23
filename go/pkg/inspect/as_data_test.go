//ff:func feature=inspect type=helper control=sequence
//ff:what TestAsData — the value when ok, nil otherwise
package inspect

import "testing"

func TestAsData(t *testing.T) {
	if got := asData(DocID{Doc: nil}, true); got == nil {
		t.Fatalf("ok: %v", got)
	}
	if got := asData(DocID{}, false); got != nil {
		t.Fatalf("not ok: %v", got)
	}
}
