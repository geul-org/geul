//ff:func feature=schema type=rule control=sequence
//ff:what TestStructure — no records, a well-formed document, a missing END
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestStructure(t *testing.T) {
	if _, fs := structure(nil, 0); rulesOf(fs) != "D1" {
		t.Fatalf("empty: %+v", fs)
	}
	ws, _ := core.ReadWords(docOf([]core.Record{metaOf(core.MetaSpecVersion, core.EncodeBytesValue(0x40, []byte("x")))}, []core.Record{core.EncodeUintValue(1)}))
	recs, _ := frame(ws)
	parts, fs := structure(recs, len(ws))
	if fs != nil || len(parts.Head) != 1 || len(parts.Body) != 1 {
		t.Fatalf("ok: %+v %+v", parts, fs)
	}
	if _, fs := structure(recs[:len(recs)-1], len(ws)); rulesOf(fs) != "D2" {
		t.Fatalf("no END: %+v", fs)
	}
}
