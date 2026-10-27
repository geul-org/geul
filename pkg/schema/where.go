//ff:type feature=schema type=model
//ff:what Where — the place of a Fact: word offset, record number and part
package schema

// Where locates a Fact. Word is the offset from the start of the document,
// Record the record number (0 is the DOC-* record) and Part the place inside
// the record ("Length", "Edge Id", "Payload" …).
type Where struct {
	Word   int
	Record int
	Part   string
}
