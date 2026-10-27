//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestApplyBind — a good name; bad layout, Subject, another document, unpublished document and a non-slug name
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestApplyBind(t *testing.T) {
	st := &listState{doc: "0xffff000000000002", docs: map[string]*docVersions{"0xffff000000000002": {}}, names: map[string]string{}}
	steps := []struct {
		r    core.Record
		want string
	}{
		{bindOf(1, subjectOf(2, 0), "sample"), ""},
		{edgeOf(core.EdgeBind, 2), "P5"},
		{bindOf(3, core.EncodeLocalRef(0), "x"), "P5"},
		{bindOf(4, subjectOf(3, 0), "x"), "P5"},
		{bindOf(5, subjectOf(2, 0), "Bad Name"), "P5"},
	}
	for i, s := range steps {
		if got := rulesOf(applyBind(st, recOf(s.r))); got != s.want {
			t.Fatalf("step %d: %q", i, got)
		}
	}
	if st.names["sample"] != "0xffff000000000002" {
		t.Fatalf("names: %v", st.names)
	}
	empty := &listState{doc: "0xffff000000000002", docs: map[string]*docVersions{}, names: map[string]string{}}
	if rulesOf(applyBind(empty, recOf(bindOf(1, subjectOf(2, 0), "sample")))) != "P5" {
		t.Fatal("unpublished: want P5")
	}
}
