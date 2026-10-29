//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what checkBytes — the schema Facts of one document's bytes as findings under a path
package cli

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/schema"
)

// checkBytes runs schema.Check (layer 1) on b. Each Fact is one finding:
// rule, severity, record and part, expected and actual.
func checkBytes(path string, b []byte) []finding {
	var out []finding
	for _, f := range schema.Check(b, schema.Options{}).Facts {
		msg := fmt.Sprintf("%s %s record %d %s: expected %s, actual %s", f.Rule, severityName(f.Severity), f.Where.Record, f.Where.Part, f.Expected, f.Actual)
		out = append(out, finding{Path: path, Message: msg})
	}
	return out
}
