//ff:func feature=notation type=decoder control=sequence
//ff:what TestFormatRecord — _ for Absent, short forms for Reference and Value, raw words otherwise
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFormatRecord(t *testing.T) {
	if formatRecord(absent()) != "_" || formatRecord(core.EncodeLocalRef(3)) != "#3" || formatRecord(core.EncodeUintValue(2)) != "UNSIGNED-INTEGER(2)" {
		t.Fatal("short forms")
	}
	if got := formatRecord(core.Record{Kind: core.KindReference, Codec: core.RefGeul << 3, Payload: []core.Word{5}}); got != "0x00900005" {
		t.Fatal(got)
	}
	if got := formatRecord(core.Record{Kind: core.KindValue, Codec: core.CodecBinaryFloat, Payload: []core.Word{0}}); len(got) < 3 || got[:2] != "0x" {
		t.Fatal(got)
	}
	if got := formatRecord(core.EncodeSpecVersion(2)); got[:2] != "0x" {
		t.Fatal(got)
	}
}
