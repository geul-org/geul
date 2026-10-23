//ff:func feature=inspect type=decoder control=sequence
//ff:what headParts — the TAG WORD, Length extension, Issuer and format number of a record as head parts
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// headParts returns the record head: the TAG WORD, the Length extension words
// (when the Length code is 1110 or 1111), and for an ISSUER codec the Issuer
// and the format number. Empty fields are left out.
func headParts(e Entry) []Part {
	iw := core.IssuerWidth(e.Record.Kind, e.Record.Codec)
	head := len(e.Words) - len(e.Record.Payload)
	ext := head - 1
	if iw > 0 {
		ext -= iw + 1
	}
	parts := []Part{{Name: "TAG", Words: e.Words[:1], Head: true}}
	if ext > 0 {
		parts = append(parts, Part{Name: "LENGTH", Words: e.Words[1 : 1+ext], Head: true})
	}
	if iw > 0 {
		parts = append(parts, Part{Name: "ISSUER", Words: e.Words[1+ext : 1+ext+iw], Head: true}, Part{Name: "FORMAT", Words: e.Words[head-1 : head], Head: true})
	}
	return parts
}
