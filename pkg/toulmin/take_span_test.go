//ff:func feature=toulmin type=decoder control=sequence
//ff:what TestTakeSpan — reads offset and length and advances; truncation fails
package toulmin

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestTakeSpan(t *testing.T) {
	pos := 0
	if s, err := TakeSpan([]core.Word{1, 2}, &pos, 1); err != nil || s != (Span{Off: 1, Len: 2}) || pos != 2 {
		t.Fatalf("got %+v %v pos=%d", s, err, pos)
	}
	pos = 0
	if _, err := TakeSpan([]core.Word{1}, &pos, 1); err == nil {
		t.Fatal("truncated: want error")
	}
}
