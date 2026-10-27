//ff:type feature=page type=model
//ff:what Ref — a link to a record by Edge Id, with that record's label
package page

// Ref links to the record with Edge Id ID; Label is that record's name when a
// LABEL gives it one.
type Ref struct {
	ID    uint64
	Label string
}
