//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what checkIssuer — version findings of every document of one Issuer
package cli

import "github.com/geul-org/geul/go/pkg/resolve"

// checkIssuer lists the Issuer's documents and checks each (checkVersions).
func checkIssuer(d resolve.Dir, issuer string) ([]finding, error) {
	k, err := d.Lookup(issuer + "/0x0000000000000000")
	if err != nil {
		return nil, err
	}
	docs, err := d.Documents(k.Issuer)
	if err != nil {
		return nil, err
	}
	var out []finding
	for _, doc := range docs {
		dk, _ := d.Lookup(issuer + "/" + doc.ID)
		out = append(out, checkVersions(d, dk.Issuer, dk.Doc)...)
	}
	return out, nil
}
