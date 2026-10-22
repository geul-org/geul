//ff:func feature=geulmd type=encoder control=sequence
//ff:what TestTidList — ids joined; an empty list keeps a lone #
package geulmd

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/toulmin"
)

func TestTidList(t *testing.T) {
	if tidList([]toulmin.TID{5, 6}) != "#5,#6" || tidList(nil) != "#" {
		t.Fatal("list mismatch")
	}
}
