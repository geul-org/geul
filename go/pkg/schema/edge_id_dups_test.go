//ff:func feature=schema type=rule control=sequence
//ff:what TestEdgeIDDups — a repeated Edge Id is reported; non-Edges are skipped
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestEdgeIDDups(t *testing.T) {
	if fs := edgeIDDups([]rec{recOf(edgeOf(0, 1)), recOf(core.EncodeUintValue(1)), recOf(edgeOf(0, 2)), recOf(edgeOf(1, 1))}); rulesOf(fs) != "E2" {
		t.Fatalf("got %+v", fs)
	}
}
