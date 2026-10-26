//ff:func feature=cli type=helper control=sequence
//ff:what TestAppContext — the data folder's specs and a Lookup that knows local key names
package cli

import (
	"testing"

	"github.com/geul-org/geul/go/internal/config"
	"github.com/geul-org/geul/go/pkg/core"
)

func TestAppContext(t *testing.T) {
	h := newHarness(t)
	h.tripleSpec()
	h.must("issuer", "new", "--name", "me")
	a := &app{env: config.Env{Getenv: func(string) string { return "" }, Cwd: h.data, ConfigDir: h.conf}}
	a.cfg, _ = config.Load(a.env)
	ctx := a.context()
	if _, ok := ctx.Spec.Of(core.Record{Kind: core.KindEdge, Codec: 0x0A}); !ok {
		t.Fatal("TRIPLE spec missing")
	}
	if k, err := ctx.Lookup("me/0x0000000000000001"); err != nil || len(k.Issuer) != 16 {
		t.Fatal(err)
	}
}
