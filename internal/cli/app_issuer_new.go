//ff:func feature=cli type=command control=sequence
//ff:what app.issuerNew — generate an Ed25519 key, save it under its local name, print the name and 16-word Issuer ID
package cli

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/geul-org/gogeul/pkg/core"
)

// keyName is a local Issuer name: a-z, 0-9 and inner '-'.
var keyName = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)

// issuerFromKey derives the Issuer of a new key; tests replace it.
var issuerFromKey = core.IssuerFromKey

// issuerNew writes `<keysDir>/<name>.pem` (PKCS #8, mode 0600, folder 0700)
// for a new key whose ID is a canonical 16-word Issuer (core.IssuerFromKey).
// An existing name is refused.
func (a *app) issuerNew(name string) error {
	if !keyName.MatchString(name) {
		return fmt.Errorf("issuer name %q: use a-z, 0-9 and -", name)
	}
	path := filepath.Join(a.keysDir(), name+".pem")
	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("issuer %q exists", name)
	}
	pub, priv, _ := ed25519.GenerateKey(rand.Reader)
	issuer, err := issuerFromKey(core.SigEd25519, pub)
	if err != nil {
		return err
	}
	der, _ := x509.MarshalPKCS8PrivateKey(priv)
	if err := os.MkdirAll(a.keysDir(), 0o700); err != nil {
		return err
	}
	if err := os.WriteFile(path, pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: der}), 0o600); err != nil {
		return err
	}
	fmt.Fprintf(a.out, "%s %s\n", name, core.FormatID(issuer))
	return nil
}
