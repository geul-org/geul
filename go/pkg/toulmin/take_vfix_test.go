//ff:func feature=toulmin type=decoder control=sequence
//ff:what TestTakeVFix — reads a SIGNED-DECIMAL as VFix; bad records and wrong codecs fail
package toulmin

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestTakeVFix(t *testing.T) {
	pos := 0
	if v, err := TakeVFix([]core.Word{0x0990, 0xFFFF, 0xFFFF, 8}, &pos); err != nil || *v != 8000 {
		t.Fatalf("got %v %v", v, err)
	}
	pos = 0
	if _, err := TakeVFix([]core.Word{0x1000}, &pos); err == nil {
		t.Fatal("non-Value: want error")
	}
	pos = 0
	if _, err := TakeVFix([]core.Word{0x0882, 7}, &pos); err == nil {
		t.Fatal("wrong codec: want error")
	}
}
