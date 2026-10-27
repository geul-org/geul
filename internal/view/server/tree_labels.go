//ff:func feature=server type=helper control=iteration dimension=1
//ff:what treeLabels — the names to show in a listing: Issuer names, plus document names inside an Issuer
package server

import "strings"

// treeLabels maps IDs to names for the listing of the directory at URL path p:
// Issuer IDs to their issuers.geul names and, inside an Issuer, its document
// IDs to their documents.geul names.
func treeLabels(root, p string) map[string]string {
	labels := issuerLabels(root)
	segs := strings.Split(strings.Trim(strings.TrimPrefix(p, wellKnown), "/"), "/")
	for id, n := range invertNames(documentNames(root, segs[0])) {
		labels[id] = n
	}
	return labels
}
