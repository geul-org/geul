//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what hashesOf — the content hashes a document's versions already have in the archive
package cli

import "github.com/geul-org/gogeul/pkg/resolve"

// hashesOf is the set of content hashes in the version list of the document
// b's DOC-ID names; empty when b has no DOC-ID or the document no list yet.
func hashesOf(d resolve.Dir, b []byte) map[[32]byte]bool {
	out := map[[32]byte]bool{}
	issuer, doc, err := docIDOf(b)
	if err != nil {
		return out
	}
	h, _ := d.Log(issuer, doc)
	for _, e := range h.Entries {
		out[e.Hash] = true
	}
	return out
}
