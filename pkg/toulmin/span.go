//ff:type feature=toulmin type=model
//ff:what Span — a source anchor: byte offset and byte length into the Source Text
package toulmin

// Span is a TANGEUL source anchor: a byte offset and a byte length into the
// document's single Source Text (GEUL t002 TOULMIN SPAN). On the wire each
// field is written at the SPAN width the SOURCE record declares.
type Span struct {
	Off uint64
	Len uint64
}
