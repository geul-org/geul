//ff:func feature=toulmin type=decoder control=sequence
//ff:what TestExpectEnd — consumed payloads pass; leftover words fail
package toulmin

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestExpectEnd(t *testing.T) {
	if err := ExpectEnd([]core.Word{1}, 1, "x"); err != nil {
		t.Fatal(err)
	}
	if err := ExpectEnd([]core.Word{1, 2}, 1, "x"); err == nil {
		t.Fatal("leftover: want error")
	}
}
