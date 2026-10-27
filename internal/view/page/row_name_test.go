//ff:func feature=page type=helper control=sequence
//ff:what TestRowName — the spec name, or kind and codec when there is none
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestRowName(t *testing.T) {
	if got := rowName(inspect.Entry{Name: "LABEL"}); got != "LABEL" {
		t.Fatalf("named: %q", got)
	}
	if got := rowName(inspect.Entry{Record: core.Record{Kind: core.KindEdge, Codec: 0x66}}); got != "EDGE 0x66" {
		t.Fatalf("unnamed: %q", got)
	}
}
