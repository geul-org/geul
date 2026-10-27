//ff:func feature=notation type=decoder control=sequence
//ff:what TestPartTokens — a generic payload of records splits; one with a raw-only record and declared parts stay one token
package notation

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestPartTokens(t *testing.T) {
	if got := partTokens(inspect.Part{Name: "payload", Words: []core.Word{0x0088, 1, 0x0088, 2}}); !slices.Equal(got, []string{"#1", "#2"}) {
		t.Fatal(got)
	}
	if got := partTokens(inspect.Part{Name: "payload", Words: []core.Word{0x0088, 1, 0x0001}}); !slices.Equal(got, []string{"0x008800010001"}) {
		t.Fatal(got)
	}
	if got := partTokens(inspect.Part{Name: "Rank", Words: []core.Word{0x0088, 1}}); !slices.Equal(got, []string{"0x00880001"}) {
		t.Fatal(got)
	}
}
