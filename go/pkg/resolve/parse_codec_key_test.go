//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what TestParseCodecKey — addresses read back; unknown Kinds, bad bits and wrong widths fail
package resolve

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestParseCodecKey(t *testing.T) {
	for _, w := range []core.Word{core.TargetWord(core.KindReference, core.CodecLocalIdentified), core.TargetWord(core.KindEdge, core.EdgeDeclare)} {
		k := codecKey(w)
		if got, ok := parseCodecKey(k.Issuer[len("proposal/"):], k.Doc); !ok || got != w {
			t.Fatalf("%04X: %04X %v", w, got, ok)
		}
	}
	for _, c := range [][2]string{{"node", "0011000"}, {"edge", "x"}, {"edge", "011000"}, {"reference", "0011000"}} {
		if _, ok := parseCodecKey(c[0], c[1]); ok {
			t.Fatalf("%v: want false", c)
		}
	}
}
