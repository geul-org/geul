//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckTarget — minimal passes, empty and leading zero fail R2
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestCheckTarget(t *testing.T) {
	if checkTarget(rec{}, []core.Word{5}) != nil || rulesOf(checkTarget(rec{}, nil)) != "R2" || rulesOf(checkTarget(rec{}, []core.Word{0, 5})) != "R2" {
		t.Fatal("wrong")
	}
}
