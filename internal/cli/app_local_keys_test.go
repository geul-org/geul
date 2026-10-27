//ff:func feature=cli type=helper control=sequence
//ff:what TestAppLocalKeys — none without a folder; other files skipped; a broken key and a keys path that is a file fails
package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/internal/config"
)

func TestAppLocalKeys(t *testing.T) {
	conf := t.TempDir()
	a := &app{env: config.Env{ConfigDir: conf}}
	if ks, err := a.localKeys(); err != nil || ks != nil {
		t.Fatal(err)
	}
	h := newHarness(t)
	h.conf = conf
	h.must("issuer", "new", "--name", "me")
	os.WriteFile(filepath.Join(conf, "geul", "keys", "README"), []byte("x"), 0o600)
	if ks, err := a.localKeys(); err != nil || len(ks) != 1 {
		t.Fatalf("%d %v", len(ks), err)
	}
	os.WriteFile(filepath.Join(conf, "geul", "keys", "broken.pem"), []byte("x"), 0o600)
	if _, err := a.localKeys(); err == nil {
		t.Fatal("broken key: want error")
	}
	blocked := t.TempDir()
	os.MkdirAll(filepath.Join(blocked, "geul"), 0o755)
	os.WriteFile(filepath.Join(blocked, "geul", "keys"), nil, 0o644)
	if _, err := (&app{env: config.Env{ConfigDir: blocked}}).localKeys(); err == nil {
		t.Fatal("keys is a file: want error")
	}
}
