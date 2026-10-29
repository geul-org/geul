//ff:func feature=cli type=helper control=sequence
//ff:what TestAppWriteOptions — the flag wins over GEULCOMPRESS; neither follows the set (nil codec); bad values fail; verification stays on
package cli

import (
	"testing"

	"github.com/geul-org/gogeul/internal/config"
	"github.com/geul-org/gogeul/pkg/pack"
)

func TestAppWriteOptions(t *testing.T) {
	env := map[string]string{}
	a := &app{}
	setEnv := func() {
		a.cfg, _ = config.Load(config.Env{Getenv: func(k string) string { return env[k] }, Cwd: t.TempDir(), ConfigDir: t.TempDir()})
	}
	setEnv()
	if o, err := a.writeOptions(""); err != nil || o.Codec != nil || !o.Verify {
		t.Fatal(o, err)
	}
	env["GEULCOMPRESS"] = "zstd:5"
	setEnv()
	if o, err := a.writeOptions(""); err != nil || o.Codec.(*pack.Zstd).Level != 5 {
		t.Fatal(o, err)
	}
	if o, err := a.writeOptions("none"); err != nil || o.Codec != (pack.Stored{}) {
		t.Fatal(o, err)
	}
	if _, err := a.writeOptions("lz4"); err == nil {
		t.Fatal("bad flag")
	}
}
