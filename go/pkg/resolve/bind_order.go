//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what bindOrder — the names a name index binds to each document ID, in BIND order
package resolve

import "github.com/geul-org/geul/go/pkg/core"

// bindOrder maps each document ID (canonical hex) a BIND of body names to its
// names in record order. Records that are not document BINDs are skipped.
func bindOrder(body []core.Record) map[string][]string {
	out := map[string][]string{}
	for _, r := range body {
		b, err := core.DecodeBind(r)
		if err != nil {
			continue
		}
		if k, ok := docOf(b.Subject); ok {
			out[k.Doc] = append(out[k.Doc], string(b.Name))
		}
	}
	return out
}
