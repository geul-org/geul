//ff:func feature=page type=builder control=iteration dimension=1
//ff:what hexParts — a record's words as head and payload HexParts from the library layout
package page

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// hexParts lays out e's words (inspect.Layout) and returns the head parts (TAG
// WORD, Length extension, Issuer, format number) and the payload parts. The
// record's TAG names a profile codec (such as RULE-DECLARE) by the record's
// name, the format number of an ISSUER record is noted with that name, and a
// Reference record's payload is split into its named fields.
func hexParts(e inspect.Entry) (head, payload []HexPart) {
	for _, p := range inspect.Layout(e) {
		h := hexPartOf(p)
		if p.Name == "TAG" {
			h.Lines[0].Note = tagBits(p.Words[0], e.Name)
		}
		if p.Name == "payload" && e.Record.Kind == core.KindReference {
			h.Lines = refPayloadLines(e.Record)
		}
		if p.Name == "FORMAT" {
			h.Lines[0].Note = e.Name
		}
		if p.Head {
			head = append(head, h)
			continue
		}
		payload = append(payload, h)
	}
	return head, payload
}
