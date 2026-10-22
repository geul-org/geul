//ff:func feature=toulmin type=encoder control=iteration dimension=1
//ff:what TestEncodeEdgeDst — FuncRef and REF dsts encode; missing FuncRef, unknown relation and provenance fail
package toulmin

import "testing"

func TestEncodeEdgeDst(t *testing.T) {
	if got, err := encodeEdgeDst(nil, EdgePacket{Prop: PropDo, DstFunc: &FuncRef{Alias: 3, Verb: 9}}, testWire); err != nil || len(got) != 2 {
		t.Fatalf("func: %v %v", got, err)
	}
	if got, err := encodeEdgeDst(nil, EdgePacket{Prop: PropLabel, DstTID: 4}, testWire); err != nil || got[0] != 4 {
		t.Fatalf("ref: %v %v", got, err)
	}
	for _, p := range []EdgePacket{{Prop: PropDo}, {Prop: 0x0FFF}, {Prop: PropProvenance}} {
		if _, err := encodeEdgeDst(nil, p, testWire); err == nil {
			t.Fatalf("0x%04X: want error", uint16(p.Prop))
		}
	}
}
