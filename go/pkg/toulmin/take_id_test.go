//ff:func feature=toulmin type=decoder control=sequence
//ff:what TestTakeID — reads a field and advances; truncation fails
package toulmin

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestTakeID(t *testing.T) {
	pos := 1
	if v, err := TakeID([]core.Word{9, 1, 2}, &pos, 2); err != nil || v != 0x10002 || pos != 3 {
		t.Fatalf("got %d %v pos=%d", v, err, pos)
	}
	pos = 0
	if _, err := TakeID([]core.Word{9}, &pos, 2); err == nil {
		t.Fatal("truncated: want error")
	}
}
