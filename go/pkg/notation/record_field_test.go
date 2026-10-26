//ff:func feature=notation type=encoder control=sequence
//ff:what TestRecordField — an allowed codec, any codec when none are declared; a disallowed codec and a bad token fail
package notation

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/codecspec"
	"github.com/geul-org/geul/go/pkg/core"
)

func TestRecordField(t *testing.T) {
	qid := codecspec.Field{Codecs: []core.Word{core.TargetWord(core.KindReference, core.RefWikidataQID<<3)}}
	if _, err := recordField("Q5", qid, Context{}); err != nil {
		t.Fatal(err)
	}
	if _, err := recordField("P5", qid, Context{}); err == nil {
		t.Fatal("disallowed: want error")
	}
	if _, err := recordField("word", recordFieldAny, Context{}); err == nil {
		t.Fatal("bad token: want error")
	}
}
