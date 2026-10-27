//ff:func feature=delta type=helper control=iteration dimension=1
//ff:what words — test fixture: records encoded back to back
package delta

import "github.com/geul-org/gogeul/pkg/core"

// words encodes rs one after another, as an Edge payload after its Edge Id.
func words(rs ...core.Record) []core.Word {
	var out []core.Word
	for _, r := range rs {
		ws, _ := core.EncodeRecord(r)
		out = append(out, ws...)
	}
	return out
}
