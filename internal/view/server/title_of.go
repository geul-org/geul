//ff:func feature=server type=helper control=sequence
//ff:what titleOf — a tree path as the page title, the Issuer and document IDs shown by their names
package server

import (
	"strings"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// titleOf is the title of the page at URL path p: the path without its
// leading /, with the Issuer ID replaced by its issuers.geul name and the
// document ID (the segment after the Issuer) by its documents.geul name, where
// they have one.
func titleOf(t resolve.Store, p string) string {
	segs := strings.Split(strings.TrimPrefix(p, wellKnown), "/")
	if len(segs) > 2 {
		if n, ok := invertNames(documentNames(t, segs[0]))[segs[1]]; ok {
			segs[1] = n
		}
	}
	if n, ok := issuerLabels(t)[segs[0]]; ok {
		segs[0] = n
	}
	return strings.TrimPrefix(wellKnown, "/") + strings.Join(segs, "/")
}
