//ff:func feature=toulmin type=helper control=iteration dimension=1
//ff:what TestControlRoundTrip — SECTION, CASE-BEGIN, CASE-END and PROSE records round-trip
package toulmin

import (
	"reflect"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// TestControlRoundTrip checks each framing control survives encode/decode at 1-
// and 2-word SPAN widths, including a case REF past one word.
func TestControlRoundTrip(t *testing.T) {
	s := Span{Off: 9, Len: 3}
	wide := Wire{Span: 2}
	cases := []struct {
		w  Wire
		in ControlPacket
	}{
		{testWire, ControlPacket{ID: 1, Code: ControlSection, Arg: uint64(SectionCases)}},
		{testWire, ControlPacket{ID: 2, Code: ControlCaseBegin, Arg: 7}},
		{testWire, ControlPacket{ID: 3, Code: ControlCaseEnd, Arg: 0xFFFF}},
		{testWire, ControlPacket{ID: 4, Code: ControlProse, Span: &s}},
		{wide, ControlPacket{ID: 5, Code: ControlCaseEnd, Arg: 0x1FFFE}},
		{wide, ControlPacket{ID: 6, Code: ControlProse, Span: &s}},
	}
	for i, c := range cases {
		ws, err := EncodeControl(c.in, c.w)
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}
		r, _, _ := core.DecodeRecord(ws)
		if got, err := DecodeControl(r, c.w); err != nil || !reflect.DeepEqual(got, c.in) {
			t.Fatalf("%d: got %+v %v want %+v", i, got, err, c.in)
		}
	}
}
