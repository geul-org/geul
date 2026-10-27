//ff:func feature=schema type=rule control=sequence
//ff:what TestMinimalUnsigned — empty, zero, leading zero
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestMinimalUnsigned(t *testing.T) {
	if minimalUnsigned(nil) || !minimalUnsigned([]core.Word{0}) || !minimalUnsigned([]core.Word{1, 0}) || minimalUnsigned([]core.Word{0, 1}) {
		t.Fatal("wrong")
	}
}
