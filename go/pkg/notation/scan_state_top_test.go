//ff:func feature=notation type=encoder control=sequence
//ff:what TestScanStateTop — only the empty state is top
package notation

import "testing"

func TestScanStateTop(t *testing.T) {
	if !(&scanState{}).top() || (&scanState{quote: true}).top() || (&scanState{angle: true}).top() || (&scanState{depth: 1}).top() {
		t.Fatal("top")
	}
}
