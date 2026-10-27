//ff:func feature=page type=helper control=sequence
//ff:what TestEdgeNodes — the Subject of a BIND, also when it names an Edge; nothing from a non-Edge
package page

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestEdgeNodes(t *testing.T) {
	if ns := edgeNodes(bindEntry(1, 7, "a")); len(ns) != 1 || ns[0].Name != "LOCAL" || ns[0].Record.Payload[0] != 7 {
		t.Fatalf("got %+v", ns)
	}
	if ns := edgeNodes(inspectEntry(core.EncodeUintValue(5), nil)); ns != nil {
		t.Fatalf("value: %+v", ns)
	}
}
