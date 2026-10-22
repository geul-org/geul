//ff:func feature=inspect type=helper control=iteration dimension=1
//ff:what TestIDOf — each TOULMIN packet's Edge Id; 0 for other content
package inspect

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/toulmin"
)

func TestIDOf(t *testing.T) {
	cases := map[uint64]any{
		1: Source{TextPacket: toulmin.TextPacket{TID: 1}},
		2: toulmin.TextPacket{TID: 2},
		3: toulmin.ControlPacket{ID: 3},
		4: toulmin.NodePacket{TID: 4},
		5: toulmin.EdgePacket{ID: 5},
		0: Value{},
	}
	for want, data := range cases {
		if got := idOf(data); got != want {
			t.Fatalf("%T: got %d want %d", data, got, want)
		}
	}
}
