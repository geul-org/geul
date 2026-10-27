//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what checkVersions — a document's version list and each live version read back (hash, diff chain) as findings
package cli

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/resolve"
)

// checkVersions reads the version list (resolve.Dir.Log) and every version
// that is not deleted (resolve.Dir.At, which rebuilds diffs and checks the
// content hash). Each failure is one finding.
func checkVersions(d resolve.Dir, issuer, doc []core.Word) []finding {
	path := core.FormatID(issuer) + "/" + core.FormatID(doc)
	h, err := d.Log(issuer, doc)
	if err != nil {
		return []finding{{Path: path, Message: "version list: " + err.Error()}}
	}
	var out []finding
	for _, e := range h.Entries {
		if e.Deleted {
			continue
		}
		if _, err := d.At(issuer, doc, e.Version); err != nil {
			out = append(out, finding{Path: fmt.Sprintf("%s@%d", path, e.Version), Message: err.Error()})
		}
	}
	return out
}
