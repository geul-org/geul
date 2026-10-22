//ff:func feature=toulmin type=decoder control=sequence
//ff:what TestPropFromCodec — DO-ONCE implies once; a node codec is not a relation
package toulmin

import "testing"

func TestPropFromCodec(t *testing.T) {
	if p, once, ok := PropFromCodec(CodecRef{Codec: CodecDoOnce}); !ok || p != PropDo || !once {
		t.Fatalf("do-once: %d %v %v", p, once, ok)
	}
	if _, _, ok := PropFromCodec(CodecRef{Codec: CodecCaseDeclare}); ok {
		t.Fatal("case-declare is not a relation")
	}
}
