//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestRecordRun — a run of embedded records passes; empty, undecodable, implausible and trailing words fail
package inspect

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestRecordRun(t *testing.T) {
	if !recordRun([]core.Word{0x0088, 7, 0x0940, 0x6d61, 0x696e}) {
		t.Fatal("LOCAL + UTF8-TEXT: want true")
	}
	for i, ws := range [][]core.Word{nil, {0xE000}, {0x0009}, {0x0088, 7, 0x0009}, {0x0900}} {
		if recordRun(ws) {
			t.Fatalf("case %d: want false", i)
		}
	}
}
