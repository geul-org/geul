//ff:func feature=resolve type=helper control=sequence
//ff:what codecDocID — the GEUL Issuer document ID of a codec spec document: 0x00000000 0001 <target word>
package resolve

import "github.com/geul-org/geul/go/pkg/core"

// codecDocID returns the 4-word document ID a codec spec document has under
// the GEUL Issuer 0x0000: 0x0000 0x0000 0x0001 then the Kind·Codec target word.
// The 0x0001 word keeps the range clear of document ID 0, which the target word
// of Reference codec 0000 would otherwise give.
func codecDocID(target core.Word) []core.Word {
	return []core.Word{0, 0, 1, target}
}
