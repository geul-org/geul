//ff:func feature=schema type=rule control=sequence
//ff:what TestListRecord — Edge Id order, BIND and PUBLICATION dispatch, foreign records
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestListRecord(t *testing.T) {
	st := &listState{docs: map[string]*docVersions{}, names: map[string]string{}, issuer: "0x0000"}
	pub := edgeOf(core.EdgePublication, 1, append([]core.Word{0xffff, 0, 0, 2, 0, 1}, make([]core.Word, 16)...)...)
	if fs := listRecord(st, 0, recOf(pub)); fs != nil {
		t.Fatalf("publication: %+v", fs)
	}
	if fs := listRecord(st, 1, recOf(bindOf(2, subjectOf(2, 0), "sample"))); fs != nil {
		t.Fatalf("bind: %+v", fs)
	}
	if fs := listRecord(st, 5, recOf(bindOf(2, subjectOf(2, 0), "sample"))); rulesOf(fs) != "P2" {
		t.Fatalf("order: %+v", fs)
	}
	if fs := listRecord(st, 3, recOf(core.EncodeUintValue(1))); rulesOf(fs) != "P1" {
		t.Fatalf("foreign: %+v", fs)
	}
}
