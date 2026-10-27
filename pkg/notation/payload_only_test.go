//ff:func feature=notation type=decoder control=sequence
//ff:what TestPayloadOnly — head parts dropped, payload parts kept in order
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestPayloadOnly(t *testing.T) {
	got := payloadOnly([]inspect.Part{{Name: "tag", Head: true}, {Name: "a"}, {Name: "b"}})
	if len(got) != 2 || got[0].Name != "a" {
		t.Fatalf("%+v", got)
	}
}
