//ff:func feature=cli type=helper control=sequence
//ff:what TestCheckFile — a valid document has no findings; a broken one names rule and severity; a missing file fails
package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestCheckFile(t *testing.T) {
	dir := t.TempDir()
	ok, _ := core.EncodeDocument([]core.Record{core.EncodeSpecVersion(2)}, nil)
	good := filepath.Join(dir, "g.geul")
	os.WriteFile(good, ok, 0o644)
	if fs, err := checkFile(good); err != nil || len(fs) != 0 {
		t.Fatalf("%+v %v", fs, err)
	}
	bad := filepath.Join(dir, "b.geul")
	os.WriteFile(bad, []byte{0, 5}, 0o644)
	if fs, _ := checkFile(bad); len(fs) == 0 || !strings.Contains(fs[0].Message, "error") {
		t.Fatalf("%+v", fs)
	}
	if _, err := checkFile(filepath.Join(dir, "none")); err == nil {
		t.Fatal("missing: want error")
	}
}
