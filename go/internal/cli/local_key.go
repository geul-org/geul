//ff:type feature=cli type=model
//ff:what localKey — a self-issued Issuer whose private key this machine holds: its local name, 16-word ID and key
package cli

import (
	"crypto/ed25519"

	"github.com/geul-org/geul/go/pkg/core"
)

// localKey is one key file of keysDir: Name is the file name without `.pem`,
// Issuer the ID computed from the public key (core.IssuerFromKey).
type localKey struct {
	Name   string
	Issuer []core.Word
	Key    ed25519.PrivateKey
}
