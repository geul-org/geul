//ff:func feature=toulmin type=decoder control=iteration dimension=1
//ff:what TestDecodeClaim — a full Claim decodes; each missing field is an error
package toulmin

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestDecodeClaim(t *testing.T) {
	pos := 0
	c, err := decodeClaim([]core.Word{4, 7, 8, 1}, &pos, ObjTID, testWire)
	if err != nil || c.Subject != 4 || c.Predicate != 7 || c.ObjTID != 8 || c.Mode != 1 || pos != 4 {
		t.Fatalf("got %+v %v pos=%d", c, err, pos)
	}
	for _, ws := range [][]core.Word{{}, {4}, {4, 7}, {4, 7, 8}} {
		p := 0
		if _, err := decodeClaim(ws, &p, ObjTID, testWire); err == nil {
			t.Fatalf("%v: want error", ws)
		}
	}
}
