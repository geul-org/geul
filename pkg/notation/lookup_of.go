//ff:func feature=notation type=encoder control=selection
//ff:what lookupOf — the context's Lookup, or one that reads canonical IDs only
package notation

import "github.com/geul-org/gogeul/pkg/resolve"

// lookupOf returns ctx.Lookup, or resolve.Dir{}.Lookup, which needs no
// archive for IDs and finds no names.
func lookupOf(ctx Context) Lookup {
	switch {
	case ctx.Lookup != nil:
		return ctx.Lookup
	default:
		return resolve.Dir{}.Lookup
	}
}
