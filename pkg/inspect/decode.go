//ff:func feature=inspect type=decoder control=sequence
//ff:what Decode — interpret a .geul byte string as one Document
package inspect

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/toulmin"
)

// Decode reads b as exactly one GEUL document and interprets every HEAD and
// BODY record. BODY records are decoded as TOULMIN records only when PROFILE
// names the TANGEUL profile this build reads. Framing errors are returned;
// a record that does not decode is kept as a bare Entry instead. Every Edge,
// a HEAD Edge (document SOURCE, its QUALIFIER) included, carries its Edge Id.
func Decode(b []byte) (Document, error) {
	ws, err := core.ReadWords(b)
	if err != nil {
		return Document{}, err
	}
	d, err := core.SplitDocument(ws)
	if err != nil {
		return Document{}, err
	}
	head, profile := decodeHead(d.Head)
	tg := toulmin.CheckProfile(profile) == nil
	body := decodeBody(d.Body, tg)
	edgeIDs(head)
	edgeIDs(body)
	return Document{Head: head, Body: body, Profile: profile, Toulmin: tg}, nil
}
