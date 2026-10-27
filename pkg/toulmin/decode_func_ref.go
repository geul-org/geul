//ff:func feature=toulmin type=decoder control=sequence
//ff:what decodeFuncRef — read a FuncRef (Alias REF, VerbCode) at *pos
package toulmin

import "github.com/geul-org/gogeul/pkg/core"

// decodeFuncRef reads a FuncRef at ws[*pos] (Alias REF then VerbCode) and
// advances *pos. Label stays empty — the stream restores it from the LABEL.
func decodeFuncRef(ws []core.Word, pos *int, w Wire) (FuncRef, error) {
	alias, err := TakeID(ws, pos)
	if err != nil {
		return FuncRef{}, err
	}
	verb, err := TakeWord(ws, pos, "FuncRef VerbCode")
	if err != nil {
		return FuncRef{}, err
	}
	return FuncRef{Alias: alias, Verb: uint16(verb)}, nil
}
