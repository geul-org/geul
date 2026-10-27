//ff:type feature=page type=model
//ff:what Field — one shown field of a record: key, text and the records it links to
package page

// Field is one field of a record: its key, its text and the records it
// references, shown as links.
type Field struct {
	Key  string
	Text string
	Refs []Ref
}
