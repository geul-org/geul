//ff:func feature=cli type=helper control=sequence
//ff:what TestAppKeysDir — keys under the config folder, never the data folder
package cli

import (
	"testing"

	"github.com/geul-org/gogeul/internal/config"
)

func TestAppKeysDir(t *testing.T) {
	a := &app{env: config.Env{ConfigDir: "/c"}}
	if a.keysDir() != "/c/geul/keys" {
		t.Fatal(a.keysDir())
	}
}
