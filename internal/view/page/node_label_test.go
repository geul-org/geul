//ff:func feature=page type=helper control=sequence
//ff:what TestNodeLabel — the LABEL name, or the record name and Edge Id
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestNodeLabel(t *testing.T) {
	e := inspect.Entry{ID: 4, Name: "RULE-DECLARE"}
	if got := nodeLabel(e, map[uint64]string{4: "규칙"}); got != "규칙" {
		t.Fatalf("labelled: %q", got)
	}
	if got := nodeLabel(e, nil); got != "RULE-DECLARE #4" {
		t.Fatalf("unlabelled: %q", got)
	}
}
