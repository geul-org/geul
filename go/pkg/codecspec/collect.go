//ff:func feature=codecspec type=decoder control=iteration dimension=1
//ff:what Collect — the Spec a document's BODY declares
package codecspec

import "github.com/geul-org/geul/go/pkg/core"

// Collect reads the codec specs in body: each DECLARE of a direct codec with
// the FIELDs that name it, in document order, and FIELD and ENUM names from
// BIND. ISSUER targets are skipped; the first DECLARE of a codec wins.
func Collect(body []core.Record) Spec {
	names := bindNames(body)
	fields := fieldDecls(body, names)
	addEnums(body, fields, names)
	s := Spec{}
	for _, r := range body {
		d, err := core.DecodeDeclare(r)
		if err != nil || len(d.Issuer) > 0 || s[d.Target] != nil {
			continue
		}
		s[d.Target] = declaredFields(d.EdgeID, body, fields)
	}
	return s
}
