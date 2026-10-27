//ff:func feature=cli type=helper control=sequence
//ff:what TestSaveDraft — a valid draft is written; one breaking the spec is not; an unwritable path fails
package cli

import (
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestSaveDraft(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "d.geul")
	if err := saveDraft(p, draft{head: []core.Record{core.EncodeSpecVersion(2)}}); err != nil {
		t.Fatal(err)
	}
	broken, _ := readDraft(brokenDraft(t, dir))
	if err := saveDraft(p, broken); err == nil {
		t.Fatal("breaks the spec: want error")
	}
	if err := saveDraft(dir, draft{}); err == nil {
		t.Fatal("folder: want error")
	}
}
