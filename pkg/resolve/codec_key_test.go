//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what TestCodecKey — every Kind's codec spec document is the GEUL Issuer document of its widened target word
package resolve

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestCodecKey(t *testing.T) {
	for want, w := range map[string]core.Word{
		"0x0000/0x0000000000000040": core.TargetWord(core.KindReference, core.RefExternalPointer1Word<<3),
		"0x0000/0x0000000000000008": core.TargetWord(core.KindReference, core.CodecLocalIdentified),
		"0x0000/0x0000000000001018": core.TargetWord(core.KindEdge, core.EdgePublication),
		"0x0000/0x0000000000000822": core.TargetWord(core.KindValue, core.CodecPosixSeconds),
		"0x0000/0x0000000000001830": core.TargetWord(core.KindMeta, 0x30),
	} {
		if k := codecKey(w); k.Issuer+"/"+k.Doc != want {
			t.Fatalf("%04X: %+v", w, k)
		}
	}
}
