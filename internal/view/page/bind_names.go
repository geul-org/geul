//ff:func feature=page type=builder control=iteration dimension=1
//ff:what bindNames — the name each BIND gives its Subject, keyed by the Subject's words
package page

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// bindNames maps each BIND Subject in body, keyed by its words as nodeIndex
// keys them, to the bound name. A later BIND of the same Subject wins (the
// last BIND is the representative name). The Subject's words are the record
// after the Edge Id, which DecodeBind has already checked.
func bindNames(body []inspect.Entry) map[string]string {
	names := map[string]string{}
	for _, e := range body {
		b, err := core.DecodeBind(e.Record)
		if err != nil {
			continue
		}
		p := e.Record.Payload
		_, n, _ := core.DecodeRecord(p)
		_, m, _ := core.DecodeRecord(p[n:])
		names[fmt.Sprint(p[n:n+m])] = string(b.Name)
	}
	return names
}
