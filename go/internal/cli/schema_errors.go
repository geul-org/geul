//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what schemaErrors — an error listing a document's schema Facts of severity error; nil when it has none
package cli

import (
	"fmt"
	"strings"

	"github.com/geul-org/geul/go/pkg/schema"
)

// schemaErrors runs schema.Check on b and fails on Error Facts (Warn Facts
// pass), naming each.
func schemaErrors(b []byte) error {
	var msgs []string
	for _, f := range schema.Check(b, schema.Options{}).Facts {
		if f.Severity == schema.Error {
			msgs = append(msgs, fmt.Sprintf("%s record %d %s: expected %s, actual %s", f.Rule, f.Where.Record, f.Where.Part, f.Expected, f.Actual))
		}
	}
	if len(msgs) == 0 {
		return nil
	}
	return fmt.Errorf("the document breaks the spec:\n  %s", strings.Join(msgs, "\n  "))
}
