//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what SplitDocument — parse a whole input as exactly one GEUL DOC into HEAD and BODY records
package core

import "errors"

// SplitDocument parses ws as exactly one GEUL document (GEUL t002 OVERVIEW 구조):
// a DOC-* record whose length covers the whole input, the HEAD Meta records, an
// optional BODY, the BODY records, and a closing END that ends the input. A
// missing END (truncated document) and any trailing record are errors.
func SplitDocument(ws []Word) (Document, error) {
	open, n, err := DecodeRecord(ws)
	if err != nil {
		return Document{}, err
	}
	if err := checkDocOpen(open, len(ws)); err != nil {
		return Document{}, err
	}
	st := &docState{}
	for i := n; i < len(ws); {
		r, m, err := DecodeRecord(ws[i:])
		if err != nil {
			return Document{}, err
		}
		if err := placeDocRecord(st, r); err != nil {
			return Document{}, err
		}
		i += m
	}
	if !st.ended {
		return Document{}, errors.New("geul/core: document has no END (truncated)")
	}
	return st.doc, nil
}
