//ff:func feature=delta type=helper control=sequence
//ff:what TestIndexOfRecord — the first equal record; none is -1
package delta

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestIndexOfRecord(t *testing.T) {
	rs := []core.Record{core.EncodeSpecVersion(1), core.EncodeSpecVersion(2), core.EncodeSpecVersion(2)}
	if indexOfRecord(rs, core.EncodeSpecVersion(2)) != 1 || indexOfRecord(rs, core.EncodeSpecVersion(3)) != -1 {
		t.Fatal("index")
	}
}
