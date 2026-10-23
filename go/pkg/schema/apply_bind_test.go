//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestApplyBind — a good name; bad layout, Subject, unpublished document, non-slug name and a name taken by another document
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestApplyBind(t *testing.T) {
	st := &listState{docs: map[string]*docVersions{"0xffff000000000002": {}, "0xffff000000000003": {}}, names: map[string]string{}, issuer: "0x0000"}
	steps := []struct {
		r    core.Record
		want string
	}{
		{bindOf(1, subjectOf(2, 0), "sample"), ""},
		{bindOf(2, subjectOf(2, 0), "sample"), ""},
		{edgeOf(core.EdgeBind, 3), "P5"},
		{bindOf(4, core.EncodeLocalRef(0), "x"), "P5"},
		{bindOf(5, subjectOf(9, 0), "x"), "P5"},
		{bindOf(6, subjectOf(3, 0), "Bad Name"), "P5"},
		{bindOf(7, subjectOf(3, 0), "sample"), "P5"},
	}
	for i, s := range steps {
		if got := rulesOf(applyBind(st, recOf(s.r))); got != s.want {
			t.Fatalf("step %d: %q", i, got)
		}
	}
}
