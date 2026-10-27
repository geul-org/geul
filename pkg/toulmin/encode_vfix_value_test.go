//ff:func feature=toulmin type=encoder control=sequence
//ff:what TestEncodeVFixValue — 0.8 encodes as SIGNED-DECIMAL C=8, E=-1
package toulmin

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestEncodeVFixValue(t *testing.T) {
	e, c, err := core.DecodeDecimalValue(EncodeVFixValue(8000))
	if err != nil || e != -1 || c != 8 {
		t.Fatalf("0.8: %d %d %v", e, c, err)
	}
}
