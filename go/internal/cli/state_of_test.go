//ff:func feature=cli type=helper control=sequence
//ff:what TestStateOf — deleted wins over deprecated; live has no suffix
package cli

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/resolve"
)

func TestStateOf(t *testing.T) {
	if stateOf(resolve.VersionEntry{Deleted: true, Deprecated: true}) != "\tdeleted" || stateOf(resolve.VersionEntry{Deprecated: true}) != "\tdeprecated" || stateOf(resolve.VersionEntry{}) != "" {
		t.Fatal("stateOf")
	}
}
