//ff:func feature=server type=helper control=sequence
//ff:what namePath — a tree path with an Issuer name and a document name replaced by their IDs; false when nothing was a name
package server

import "strings"

// namePath reads the first segment of URL path p as an Issuer name
// (issuers.geul) and the segment after it as a document name
// (documents.geul), and returns p with the IDs in their place. It is false
// when neither is a known name.
func namePath(root, p string) (string, bool) {
	segs := strings.Split(strings.TrimPrefix(p, wellKnown), "/")
	issuer, named := issuerNames(root)[segs[0]]
	if named {
		segs[0] = issuer
	}
	if len(segs) > 1 {
		if doc, ok := documentNames(root, segs[0])[segs[1]]; ok {
			segs[1], named = doc, true
		}
	}
	return wellKnown + strings.Join(segs, "/"), named
}
