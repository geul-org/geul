//ff:func feature=schema type=rule control=sequence
//ff:what TestStrictify — Warns become Errors
package schema

import (
	"testing"
)

func TestStrictify(t *testing.T) {
	if fs := strictify([]Fact{{Severity: Warn}, {Severity: Error}}); fs[0].Severity != Error || fs[1].Severity != Error {
		t.Fatalf("got %+v", fs)
	}
}
