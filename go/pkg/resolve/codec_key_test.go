//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what TestCodecKey — Reference CCCC bits and 7-bit codecs of the other Kinds
package resolve

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestCodecKey(t *testing.T) {
	for want, w := range map[string]core.Word{
		"proposal/reference/0000": core.TargetWord(core.KindReference, 0),
		"proposal/reference/1000": core.TargetWord(core.KindReference, core.CodecLocalIdentified),
		"proposal/edge/0011000":   core.TargetWord(core.KindEdge, core.EdgePublication),
		"proposal/value/0100010":  core.TargetWord(core.KindValue, core.CodecPosixSeconds),
		"proposal/meta/0110000":   core.TargetWord(core.KindMeta, 0x30),
	} {
		if k := codecKey(w); k.Issuer+"/"+k.Doc != want {
			t.Fatalf("%04X: %+v", w, k)
		}
	}
}
