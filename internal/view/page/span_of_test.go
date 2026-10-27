//ff:func feature=page type=helper control=iteration dimension=1
//ff:what TestSpanOf — NODE, LINK and PROSE spans; a control without one and other content have none
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestSpanOf(t *testing.T) {
	sp := toulmin.Span{Off: 1, Len: 2}
	for _, data := range []any{toulmin.NodePacket{Span: sp}, toulmin.EdgePacket{Span: sp}, toulmin.ControlPacket{Span: &sp}} {
		if got, ok := spanOf(data); !ok || got != sp {
			t.Fatalf("%T: %+v %v", data, got, ok)
		}
	}
	if _, ok := spanOf(toulmin.ControlPacket{}); ok {
		t.Fatal("section: want none")
	}
	if _, ok := spanOf(toulmin.TextPacket{}); ok {
		t.Fatal("text: want none")
	}
}
