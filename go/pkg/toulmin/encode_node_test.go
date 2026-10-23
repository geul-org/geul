//ff:func feature=toulmin type=encoder control=iteration dimension=1
//ff:what TestEncodeNode — NODE golden bytes plus unknown type, missing-field and SPAN overflow errors
package toulmin

import (
	"reflect"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

// TestEncodeNode pins a CASE-DECLARE record (TAG 0x12E0 = Edge, 6 words, codec
// 0x60; LOCAL Edge Id; header; SPAN) and an ALIAS record under ISSUER-1WORD, and
// checks each error: unknown NodeType, every flag whose field is absent, and a
// SPAN beyond the width.
func TestEncodeNode(t *testing.T) {
	got, err := EncodeNode(NodePacket{TID: 7, Type: NodeCase, Span: Span{Off: 1, Len: 2}}, testWire)
	if want := []core.Word{0x12E0, 0x00C0, 7, 0, 1, 2}; err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("case: %04X %v want %04X", got, err, want)
	}
	got, err = EncodeNode(NodePacket{TID: 1, Type: NodeAlias}, testWire)
	if want := []core.Word{0x13FC, 0, 0xC013, 0x00C0, 1, 0, 0, 0}; err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("alias: %04X %v want %04X", got, err, want)
	}
	bad := []NodePacket{
		{Type: 0x00FF},
		{Type: NodeCase, Flags: NodeFlagFunc},
		{Type: NodeCase, Flags: NodeFlagQual},
		{Type: NodeCase, Flags: NodeFlagCheck},
		{Type: NodeCase, Flags: NodeFlagClaim},
		{Type: NodeCase, Flags: NodeFlagVerdict},
		{Type: NodeCase, Flags: NodeFlagClaim, Claim: &Claim{Kind: 3}},
		{Type: NodeCase, Span: Span{Off: 0x10000}},
	}
	for i, p := range bad {
		if _, err := EncodeNode(p, testWire); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
