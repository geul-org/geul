//ff:type feature=page type=model
//ff:what HexPart — one field of a record's words, one line per word; an embedded record keeps its own lines
package page

// HexPart is one field of a record's words for display, one HexLine per word.
// Record marks an embedded record (drawn as a nested box named Name), whose
// first line is its own TAG WORD.
type HexPart struct {
	Name   string
	Record bool
	Lines  []HexLine
}
