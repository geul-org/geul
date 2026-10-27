//ff:func feature=toulmin type=decoder control=sequence
//ff:what TestNodeFromCodec — rule codecs fix the role; non-node codecs are not ok
package toulmin

import "testing"

func TestNodeFromCodec(t *testing.T) {
	nt, role, fixed, ok := NodeFromCodec(CodecRef{Codec: CodecCounterDeclare})
	if !ok || nt != NodeRule || role != RoleCounter || !fixed {
		t.Fatalf("counter: %d %d %v %v", nt, role, fixed, ok)
	}
	if _, _, fixed, ok := NodeFromCodec(CodecRef{Codec: CodecSpecDeclare}); !ok || fixed {
		t.Fatal("spec-declare is a node without a fixed role")
	}
	if _, _, _, ok := NodeFromCodec(CodecRef{Codec: CodecAttacks}); ok {
		t.Fatal("attacks is not a node codec")
	}
}
