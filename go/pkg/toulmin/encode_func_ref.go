//ff:func feature=toulmin type=encoder control=sequence
//ff:what encodeFuncRef — append a FuncRef as the Alias REF then the VerbCode word
package toulmin

import "github.com/geul-org/geul/go/pkg/core"

// encodeFuncRef appends f to out: the Alias REF at the Edge Id width, then the
// VerbCode word. Label is not written — it rides as a TEXT + LABEL pair.
func encodeFuncRef(out []core.Word, f FuncRef, w Wire) ([]core.Word, error) {
	out, err := PutID(out, f.Alias, w.ID)
	if err != nil {
		return nil, err
	}
	return append(out, core.Word(f.Verb)), nil
}
