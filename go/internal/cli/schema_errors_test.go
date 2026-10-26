//ff:func feature=cli type=helper control=sequence
//ff:what TestSchemaErrors — a valid document is nil; error Facts are listed
package cli

import (
	"strings"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestSchemaErrors(t *testing.T) {
	ok, _ := core.EncodeDocument([]core.Record{core.EncodeSpecVersion(2)}, nil)
	if err := schemaErrors(ok); err != nil {
		t.Fatal(err)
	}
	if err := schemaErrors([]byte{0, 5}); err == nil || !strings.Contains(err.Error(), "breaks the spec") {
		t.Fatal(err)
	}
}
