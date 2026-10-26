//ff:func feature=cli type=helper control=sequence
//ff:what readKey — a local Issuer key file: PEM PKCS #8 Ed25519 private key, its Issuer computed
package cli

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/geul-org/geul/go/pkg/core"
)

// readKey reads path and returns its localKey. A file that is not a PEM
// Ed25519 private key is an error.
func readKey(path string) (localKey, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return localKey{}, err
	}
	blk, _ := pem.Decode(b)
	if blk == nil {
		return localKey{}, fmt.Errorf("%s: not a PEM file", path)
	}
	k, err := x509.ParsePKCS8PrivateKey(blk.Bytes)
	priv, ok := k.(ed25519.PrivateKey)
	if err != nil || !ok {
		return localKey{}, fmt.Errorf("%s: not an Ed25519 private key", path)
	}
	issuer, err := core.IssuerFromKey(core.SigEd25519, priv.Public().(ed25519.PublicKey))
	return localKey{Name: strings.TrimSuffix(filepath.Base(path), ".pem"), Issuer: issuer, Key: priv}, err
}
