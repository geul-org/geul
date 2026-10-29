//ff:func feature=server type=helper control=iteration dimension=1
//ff:what documentNames — document name → ID from a document space's documents.geul; empty when it is missing or broken
package server

import (
	"github.com/geul-org/gogeul/pkg/resolve"
)

// documentNames reads `<space>/documents.geul` (GEUL-ORG 문서 이름) of an
// Issuer and keeps the names of that Issuer's own documents.
func documentNames(t resolve.Store, space string) map[string]string {
	out := map[string]string{}
	b, err := t.ReadFile(space + "/documents.geul")
	if err != nil {
		return out
	}
	names, _ := resolve.ParseDocumentNames(b)
	for n, d := range names {
		if d.Issuer == space {
			out[n] = d.Doc
		}
	}
	return out
}
