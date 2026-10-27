//ff:func feature=schema type=rule control=sequence
//ff:what TestRunCodecs — every record is offered to the Codecs
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestRunCodecs(t *testing.T) {
	recs := []rec{{R: core.Record{Codec: 5}}, {R: core.Record{Codec: 5}}, {R: core.Record{Codec: 6}}}
	if fs := runCodecs(recs, []Codec{fakeCodec{codec: 5}}); rulesOf(fs) != "X1,X1" {
		t.Fatalf("got %+v", fs)
	}
}
