//ff:func feature=toulmin type=decoder control=sequence
//ff:what TestTakeWord — reads one word and advances; past the end fails
package toulmin

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestTakeWord(t *testing.T) {
	pos := 0
	if w, err := TakeWord([]core.Word{7}, &pos, "x"); err != nil || w != 7 || pos != 1 {
		t.Fatalf("got %d %v pos=%d", w, err, pos)
	}
	if _, err := TakeWord([]core.Word{7}, &pos, "x"); err == nil {
		t.Fatal("past end: want error")
	}
}
