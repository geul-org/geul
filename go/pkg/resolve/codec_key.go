//ff:func feature=resolve type=helper control=sequence
//ff:what codecKey — `0x0000/<codecDocID>`, the address of a direct codec's spec document
package resolve

import "github.com/geul-org/geul/go/pkg/core"

// codecKey is the address of the spec document of the codec in target
// (GEUL-ORG 직접 코드 코덱북): an ordinary document of the GEUL Issuer 0x0000
// whose document ID is the target word widened to 4 words (codecDocID).
func codecKey(target core.Word) DocKey {
	return DocKey{Issuer: core.FormatID([]core.Word{0}), Doc: core.FormatID(codecDocID(target))}
}
