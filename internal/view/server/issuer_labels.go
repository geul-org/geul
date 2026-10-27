//ff:func feature=server type=helper control=sequence
//ff:what issuerLabels — Issuer ID → bound name from the tree's issuers.geul; empty when it is missing or broken
package server

// issuerLabels inverts the names of issuers.geul at the top of the tree.
func issuerLabels(root string) map[string]string {
	return invertNames(issuerNames(root))
}
