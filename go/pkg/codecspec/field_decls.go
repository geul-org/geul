//ff:func feature=codecspec type=decoder control=iteration dimension=1
//ff:what fieldDecls — every FIELD in body as a named Field by Edge Id
package codecspec

import "github.com/geul-org/geul/go/pkg/core"

// fieldDecls reads every FIELD Edge of body into a Field named by its BIND.
func fieldDecls(body []core.Record, names map[uint64]string) map[uint64]*Field {
	out := map[uint64]*Field{}
	for _, r := range body {
		f, err := core.DecodeFieldDecl(r)
		if err != nil {
			continue
		}
		out[f.EdgeID] = &Field{Name: names[f.EdgeID], Width: int(f.Width), Enums: map[string]string{}}
	}
	return out
}
