//ff:func feature=page type=helper control=iteration dimension=1
//ff:what recordsText — the records in a run of words, each shown by recordText, joined by spaces
package page

import (
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// recordsText shows the records ws holds one after another (recordText),
// joined by spaces. Words left after a record that does not decode are shown
// as one hex run.
func recordsText(ws []core.Word) string {
	var parts []string
	for len(ws) > 0 {
		r, n, err := core.DecodeRecord(ws)
		if err != nil {
			parts = append(parts, core.FormatID(ws))
			break
		}
		parts = append(parts, recordText(r))
		ws = ws[n:]
	}
	return strings.Join(parts, " ")
}
