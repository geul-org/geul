//ff:func feature=cli type=helper control=sequence
//ff:what app.context — the notation context: codec specs of the data folder and name lookup through app.key
package cli

import "github.com/geul-org/geul/go/pkg/notation"

// context reads the codec specs of the data folder (none without one) for
// field layouts and resolves names with key.
func (a *app) context() notation.Context {
	return notation.Context{Spec: a.spec(), Lookup: a.key}
}
