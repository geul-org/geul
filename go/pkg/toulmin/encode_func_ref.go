//ff:func feature=toulmin type=encoder control=sequence
//ff:what encodeFuncRef — append a FuncRef as the Alias REF then the VerbCode word
package toulmin

import "github.com/geul-org/geul/go/pkg/core"

// encodeFuncRef appends f to out: the Alias REF (a LOCAL Reference), then the
// VerbCode word. Label is not written — it rides as a TEXT + LABEL pair.
func encodeFuncRef(out []core.Word, f FuncRef) []core.Word {
	return append(PutID(out, f.Alias), core.Word(f.Verb))
}
