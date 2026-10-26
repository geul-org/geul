//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what issuerKeyHead — a HEAD with ISSUER-KEY added when the Issuer is a local self-issued key
package cli

import (
	"crypto/ed25519"
	"slices"

	"github.com/geul-org/geul/go/pkg/core"
)

// issuerKeyHead appends ISSUER-KEY (meta/verification/ISSUER-KEY) to head
// when issuer is one of keys, so readers can check the 16-word ID against the
// key.
func issuerKeyHead(keys []localKey, issuer []core.Word, head []core.Record) []core.Record {
	for _, k := range keys {
		if slices.Equal(k.Issuer, issuer) {
			r, _ := core.EncodeIssuerKey(core.SigEd25519, k.Key.Public().(ed25519.PublicKey))
			return append(head, r)
		}
	}
	return head
}
