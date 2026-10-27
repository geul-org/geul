//ff:type feature=page type=model
//ff:what HexLine — one word of a record on its own line: hex, field name and note
package page

// HexLine is one displayed word. Word is its hex (or "…" for a cut run), Name
// the field it starts, Note what it means — a TAG WORD's bit fields, a format
// number's record name, a text Value's string, or how many words a cut hid.
// Tag marks a TAG WORD and Text a Note that is a decoded string (shown bold).
type HexLine struct {
	Word string
	Name string
	Note string
	Tag  bool
	Text bool
}
