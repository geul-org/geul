//ff:func feature=codecspec type=helper control=sequence
//ff:what keyOf — the Spec key of a record: its target word, Reference Quantifier cleared
package codecspec

import "github.com/geul-org/geul/go/pkg/core"

// keyOf is the Spec key of r: its Kind and Codec, with a Reference's
// Quantifier bits cleared (edge/DECLARE).
func keyOf(r core.Record) core.Word {
	c := r.Codec
	if r.Kind == core.KindReference {
		c &^= 7
	}
	return core.TargetWord(r.Kind, c)
}
