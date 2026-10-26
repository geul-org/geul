//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what addressesOf — the addresses <Issuer>/<문서> of an Issuer's documents
package cli

import "github.com/geul-org/geul/go/pkg/resolve"

// addressesOf joins issuer and each document ID.
func addressesOf(issuer string, docs []resolve.Named) []string {
	var out []string
	for _, d := range docs {
		out = append(out, issuer+"/"+d.ID)
	}
	return out
}
