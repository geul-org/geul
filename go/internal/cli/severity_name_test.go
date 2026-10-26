//ff:func feature=cli type=helper control=sequence
//ff:what TestSeverityName — error and warn
package cli

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/schema"
)

func TestSeverityName(t *testing.T) {
	if severityName(schema.Error) != "error" || severityName(schema.Warn) != "warn" {
		t.Fatal("names")
	}
}
