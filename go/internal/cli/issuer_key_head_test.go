//ff:func feature=cli type=helper control=sequence
//ff:what TestIssuerKeyHead — ISSUER-KEY added for the matching local key only
package cli

import (
	"crypto/ed25519"
	"crypto/rand"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestIssuerKeyHead(t *testing.T) {
	pub, priv, _ := ed25519.GenerateKey(rand.Reader)
	issuer, _ := core.IssuerFromKey(core.SigEd25519, pub)
	keys := []localKey{{Name: "other", Issuer: make([]core.Word, 16), Key: priv}, {Name: "me", Issuer: issuer, Key: priv}}
	head := []core.Record{core.EncodeSpecVersion(2)}
	got := issuerKeyHead(keys, issuer, head)
	if len(got) != 2 || got[1].Codec != core.MetaIssuerKey {
		t.Fatalf("%+v", got)
	}
	if got := issuerKeyHead(keys, []core.Word{0}, head); len(got) != 1 {
		t.Fatal("not local: want the head as is")
	}
}
