//ff:func feature=geulmd type=encoder control=sequence
//ff:what TestMetaValue — one known Value renders; trailing words or a bad record do not
package geulmd

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestMetaValue(t *testing.T) {
	vw, _ := core.EncodeRecord(core.EncodeUintValue(3))
	if got, ok := metaValue("X", core.Record{Payload: vw}); !ok || got != "X uint(3)" {
		t.Fatalf("got %q %v", got, ok)
	}
	if _, ok := metaValue("X", core.Record{Payload: append(vw, 0)}); ok {
		t.Fatal("trailing word: want not ok")
	}
	if _, ok := metaValue("X", core.Record{}); ok {
		t.Fatal("empty: want not ok")
	}
}
