//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what declareTarget — the target of the only DECLARE in a BODY; false for none, several or an ISSUER target
package resolve

import "github.com/geul-org/geul/go/pkg/core"

// declareTarget returns the target word of the one DECLARE in body.
func declareTarget(body []core.Record) (core.Word, bool) {
	var found []core.Declare
	for _, r := range body {
		if dc, err := core.DecodeDeclare(r); err == nil {
			found = append(found, dc)
		}
	}
	if len(found) != 1 || len(found[0].Issuer) > 0 {
		return 0, false
	}
	return found[0].Target, true
}
