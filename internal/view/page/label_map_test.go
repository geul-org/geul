//ff:func feature=page type=builder control=sequence
//ff:what TestLabelMap — a LABEL names its src with the TEXT it points at; other links do not label
package page

import "testing"

func TestLabelMap(t *testing.T) {
	labels := labelMap(fixtureDoc(t).Body)
	if len(labels) != 1 || labels[1] != "사안" {
		t.Fatalf("got %v", labels)
	}
}
