//ff:func feature=page type=helper control=sequence
//ff:what recordLines — an embedded record's words: TAG, Length extension, then payload (Reference fields named)
package page

import "github.com/geul-org/gogeul/pkg/core"

// recordLines renders an embedded record one word per line: its TAG WORD, any
// Length extension words, then its payload. A Reference payload is split into
// its named fields; any other payload is one "payload" run.
func recordLines(ws []core.Word) []HexLine {
	lines := hexLines(ws[:1], "TAG")
	r, _, err := core.DecodeRecord(ws)
	if err != nil || r.Kind != core.KindReference {
		return append(lines, hexLines(ws[1:], "payload")...)
	}
	head := len(ws) - len(r.Payload)
	lines = append(lines, hexLines(ws[1:head], "LENGTH")...)
	return append(lines, refPayloadLines(r)...)
}
