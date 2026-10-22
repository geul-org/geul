//ff:func feature=inspect type=decoder control=sequence
//ff:what TestNewEntry — words include the TAG WORD and the name is the core codec name
package inspect

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestNewEntry(t *testing.T) {
	e := newEntry(core.EncodeUintValue(5))
	if len(e.Words) != 2 || e.Words[1] != 5 || e.Name != "UNSIGNED-INTEGER" || e.Data != nil {
		t.Fatalf("got %+v", e)
	}
}
