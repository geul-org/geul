//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestReadKey — a key made by issuer new reads back; a missing file, non-PEM, a non-Ed25519 key and bad DER fail
package cli

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"
	"testing"
)

func TestReadKey(t *testing.T) {
	h := newHarness(t)
	h.must("issuer", "new", "--name", "me")
	dir := filepath.Join(h.conf, "geul", "keys")
	if k, err := readKey(filepath.Join(dir, "me.pem")); err != nil || k.Name != "me" || len(k.Issuer) != 16 {
		t.Fatalf("%+v %v", k, err)
	}
	ec, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	der, _ := x509.MarshalPKCS8PrivateKey(ec)
	files := map[string][]byte{
		"text.pem":  []byte("x"),
		"ecdsa.pem": pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: der}),
		"junk.pem":  pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: []byte{1, 2}}),
	}
	for name, b := range files {
		os.WriteFile(filepath.Join(dir, name), b, 0o600)
		if _, err := readKey(filepath.Join(dir, name)); err == nil {
			t.Errorf("%s: want error", name)
		}
	}
	if _, err := readKey(filepath.Join(dir, "none.pem")); err == nil {
		t.Fatal("missing: want error")
	}
}
