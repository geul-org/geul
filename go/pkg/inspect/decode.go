//ff:func feature=inspect type=decoder control=sequence
//ff:what Decode — interpret a .geul byte string as one Document
package inspect

import (
	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/toulmin"
)

// Decode reads b as exactly one GEUL document and interprets every HEAD and
// BODY record. BODY records are decoded as TOULMIN records only when PROFILE
// names the TANGEUL profile this build reads. Framing errors are returned;
// a record that does not decode is kept as a bare Entry instead.
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
	return Document{Head: head, Body: decodeBody(d.Body, tg), Profile: profile, Toulmin: tg}, nil
}
