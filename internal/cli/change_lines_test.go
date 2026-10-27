//ff:func feature=cli type=helper control=sequence
//ff:what TestChangeLines — + has only a new line, - only an old one, ~ both
package cli

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/delta"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestChangeLines(t *testing.T) {
	r := core.EncodeSpecVersion(2)
	if l := changeLines(delta.Change{Op: '+', New: r}, inspect.Document{}, inspect.Document{}); l.Old != "" || l.New == "" {
		t.Fatalf("%+v", l)
	}
	if l := changeLines(delta.Change{Op: '-', Old: r}, inspect.Document{}, inspect.Document{}); l.Old == "" || l.New != "" {
		t.Fatalf("%+v", l)
	}
	if l := changeLines(delta.Change{Op: '~', Old: r, New: r, ID: 3}, inspect.Document{}, inspect.Document{}); l.Old == "" || l.New == "" || l.ID != 3 {
		t.Fatalf("%+v", l)
	}
}
