//ff:func feature=schema type=decoder control=sequence
//ff:what listDoc — the document a version list lists: the document ID of its first PUBLICATION
package schema

import "github.com/geul-org/gogeul/pkg/core"

// listDoc returns, in canonical hex, the document ID of the version list's
// first record (a PUBLICATION, as isVersionList requires); "" when it does not decode.
func listDoc(body []rec) string {
	p, err := core.DecodePublication(body[0].R)
	if err != nil {
		return ""
	}
	return core.FormatID(p.Doc)
}
