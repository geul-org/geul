//ff:func feature=page type=helper control=sequence
//ff:what TestWordsNumber — one, two and four words big-endian; none is zero
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestWordsNumber(t *testing.T) {
	if wordsNumber([]core.Word{42}) != 42 || wordsNumber([]core.Word{0x12, 0xD687}) != 1234567 || wordsNumber([]core.Word{0, 0, 1, 0}) != 65536 || wordsNumber(nil) != 0 {
		t.Fatal("wrong number")
	}
}
