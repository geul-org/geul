//ff:func feature=page type=helper control=iteration dimension=1
//ff:what TestClassOf — every decoded content type has its class; SECTION is set apart from other controls
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/inspect"
	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestClassOf(t *testing.T) {
	cases := map[string]any{
		"value":   inspect.Value{},
		"source":  inspect.Source{},
		"text":    toulmin.TextPacket{},
		"section": toulmin.ControlPacket{Code: toulmin.ControlSection},
		"control": toulmin.ControlPacket{Code: toulmin.ControlProse},
		"node":    toulmin.NodePacket{},
		"link":    toulmin.EdgePacket{},
		"raw":     nil,
	}
	for want, data := range cases {
		if got := classOf(data); got != want {
			t.Fatalf("%T: got %q want %q", data, got, want)
		}
	}
}
