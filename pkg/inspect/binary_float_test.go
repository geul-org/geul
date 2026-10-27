//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestBinaryFloat — binary16, binary32 and binary64 Payloads decode; binary128 and odd widths fail
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestBinaryFloat(t *testing.T) {
	if f, err := binaryFloat([]core.Word{0x3C00}); err != nil || f != 1 {
		t.Fatalf("16: %v %v", f, err)
	}
	if f, err := binaryFloat([]core.Word{0x4049, 0x0FDB}); err != nil || float32(f) != float32(3.1415927) {
		t.Fatalf("32: %v %v", f, err)
	}
	if f, err := binaryFloat([]core.Word{0x4012, 0, 0, 0}); err != nil || f != 4.5 {
		t.Fatalf("64: %v %v", f, err)
	}
	for _, ws := range [][]core.Word{make([]core.Word, 8), make([]core.Word, 3), nil} {
		if _, err := binaryFloat(ws); err == nil {
			t.Fatalf("%d words: want error", len(ws))
		}
	}
}
