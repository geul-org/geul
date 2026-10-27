//ff:func feature=page type=helper control=sequence
//ff:what TestCodeName — a named code and an unnamed one
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestCodeName(t *testing.T) {
	if got := codeName(levelNames, toulmin.LevelFail); got != "fail" {
		t.Fatalf("named: %q", got)
	}
	if got := codeName(levelNames, toulmin.LevelCode(3)); got != "3" {
		t.Fatalf("unnamed: %q", got)
	}
}
