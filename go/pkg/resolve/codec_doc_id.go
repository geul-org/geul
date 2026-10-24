//ff:func feature=resolve type=helper control=sequence
//ff:what codecDocID — the GEUL Issuer document ID of a codec spec document: the target word widened to 4 words
package resolve

import "github.com/geul-org/geul/go/pkg/core"

// codecDocID returns the 4-word document ID a codec spec document has under
// the GEUL Issuer 0x0000: the Kind·Codec target word with three 0x0000 words
// before it (TAG 0x1234 → 0x0000000000001234). No target word is 0, since the
// Reference codec 0000 is reserved.
func codecDocID(target core.Word) []core.Word {
	return []core.Word{0, 0, 0, target}
}
