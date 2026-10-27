//ff:type feature=page type=model
//ff:what Row — one record on the page: name, Edge Id, label, fields, back-references, span and words
package page

// Row is one record. Anchor is its fragment id (`e<Edge Id>`, or `h<n>`/`r<n>`
// for records without one). Class is the record class for styling. Label is
// the name a LABEL record gives it. Back lists the records that reference it.
// HeadParts are the TAG WORD and its extensions, BodyParts the payload fields,
// both in hex (hexParts). Summary is the field texts on one line for the list
// (summaryOf). Depth is the case nesting. Edge marks a record of
// the Edge Kind.
type Row struct {
	Index     int
	Anchor    string
	Name      string
	Class     string
	ID        uint64
	Label     string
	Fields    []Field
	Summary   string
	Back      []Ref
	HasSpan   bool
	SpanOff   uint64
	SpanLen   uint64
	HeadParts []HexPart
	BodyParts []HexPart
	Depth     int
	Edge      bool
}
