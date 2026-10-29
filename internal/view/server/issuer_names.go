//ff:func feature=server type=helper control=sequence
//ff:what issuerNames — Issuer name → ID from the tree's issuers.geul; empty when it is missing or broken
package server

import (
	"github.com/geul-org/gogeul/pkg/resolve"
)

// issuerNames reads issuers.geul at the top of the tree (GEUL-ORG Issuer 이름).
func issuerNames(t resolve.Store) map[string]string {
	b, err := t.ReadFile("issuers.geul")
	if err != nil {
		return map[string]string{}
	}
	names, err := resolve.ParseIssuerNames(b)
	if err != nil {
		return map[string]string{}
	}
	return names
}
