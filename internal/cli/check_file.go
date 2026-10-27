//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what checkFile — the schema Facts of one file as findings
package cli

import (
	"fmt"
	"os"

	"github.com/geul-org/gogeul/pkg/schema"
)

// checkFile reads path and runs schema.Check (layer 1). Each Fact is one
// finding: rule, severity, record and part, expected and actual.
func checkFile(path string) ([]finding, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var out []finding
	for _, f := range schema.Check(b, schema.Options{}).Facts {
		msg := fmt.Sprintf("%s %s record %d %s: expected %s, actual %s", f.Rule, severityName(f.Severity), f.Where.Record, f.Where.Part, f.Expected, f.Actual)
		out = append(out, finding{Path: path, Message: msg})
	}
	return out, nil
}
