//ff:func feature=schema type=rule control=sequence
//ff:what TestCodecFacts — only matching Codecs run, at the record's place
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestCodecFacts(t *testing.T) {
	r := rec{Index: 2, Word: 7, R: core.Record{Codec: 5}}
	fs := codecFacts(r, []Codec{fakeCodec{codec: 5}, fakeCodec{codec: 6}})
	if rulesOf(fs) != "X1" || fs[0].Where != (Where{Word: 7, Record: 2}) {
		t.Fatalf("got %+v", fs)
	}
}
