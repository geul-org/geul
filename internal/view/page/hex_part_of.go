//ff:func feature=page type=builder control=sequence
//ff:what hexPartOf — one layout part as a HexPart: a line per word, TAG WORDs decoded
package page

import "github.com/geul-org/gogeul/pkg/inspect"

// hexPartOf renders p one word per line. The record's own TAG part and an
// embedded record's first word are TAG WORDs and carry their bit fields; an
// embedded record's remaining words are its payload (a Reference's fields
// named), and an embedded UTF8-TEXT
// Value notes its string beside the first payload word. A raw field with an
// ENUM name notes it on its first word.
func hexPartOf(p inspect.Part) HexPart {
	lines := hexLines(p.Words, p.Name)
	if p.Record {
		lines = recordLines(p.Words)
	}
	if p.Record || p.Name == "TAG" {
		lines[0].Note, lines[0].Tag = tagBits(p.Words[0], ""), true
	}
	if s, ok := textOf(p.Words); p.Record && ok && len(lines) > 1 {
		lines[1].Note, lines[1].Text = s, true
	}
	if p.Enum != "" {
		lines[0].Note = p.Enum
	}
	return HexPart{Name: p.Name, Record: p.Record, Lines: lines}
}
