//ff:func feature=core type=helper control=selection
//ff:what Kind.String — the spec name of a record kind (REFERENCE/VALUE/EDGE/META)
package core

// String returns the spec name of k (GEUL t002 README TAG WORD).
func (k Kind) String() string {
	switch k {
	case KindReference:
		return "REFERENCE"
	case KindValue:
		return "VALUE"
	case KindEdge:
		return "EDGE"
	default:
		return "META"
	}
}
