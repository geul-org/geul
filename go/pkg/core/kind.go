//ff:type feature=core type=model
//ff:what Kind — the 2-bit record kind of a GEUL TAG WORD (Reference/Value/Edge/Meta)
package core

// Kind is the 2-bit record kind in bits 12-11 of a GEUL TAG WORD
// `000K KLLL LCCC CCCC` (GEUL t002 README). A decoder tells Reference, Value,
// Edge and Meta apart from the TAG WORD alone, without a codebook.
type Kind uint8

const (
	// KindReference is a record that points at a target (00).
	KindReference Kind = 0
	// KindValue is a record that carries a value (01).
	KindValue Kind = 1
	// KindEdge is a record that composes other records (10).
	KindEdge Kind = 2
	// KindMeta is a record about the document (11).
	KindMeta Kind = 3
)
