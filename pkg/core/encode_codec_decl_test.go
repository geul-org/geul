//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeCodecDecl — the docs' LABEL Object example words
package core

import (
	"slices"
	"testing"
)

func TestEncodeCodecDecl(t *testing.T) {
	ws, _ := EncodeRecord(EncodeCodecDecl(CodecDecl{EdgeID: 9, Field: 4, Target: TargetWord(KindValue, CodecUTF8LanguageTextEven)}))
	if !slices.Equal(ws, []Word{0x129E, 0x0088, 0x0009, 0x0088, 0x0004, 0x084A}) {
		t.Fatalf("%04X", ws)
	}
}
