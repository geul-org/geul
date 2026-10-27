//ff:func feature=cli type=helper control=sequence
//ff:what TestAppPrintFindings — path: message lines, or JSON
package cli

import (
	"bytes"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/internal/config"
)

func TestAppPrintFindings(t *testing.T) {
	var b bytes.Buffer
	a := &app{out: &b}
	a.printFindings([]finding{{Path: "p", Message: "m"}})
	if b.String() != "p: m\n" {
		t.Fatal(b.String())
	}
	b.Reset()
	a.cfg, _ = config.Load(config.Env{Flags: map[string]string{"GEULFORMAT": "json"}, Getenv: func(string) string { return "" }, Cwd: t.TempDir(), ConfigDir: t.TempDir()})
	a.printFindings([]finding{{Path: "p", Message: "m"}})
	if !strings.Contains(b.String(), `"path": "p"`) {
		t.Fatal(b.String())
	}
}
