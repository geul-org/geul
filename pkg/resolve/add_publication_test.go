//ff:func feature=resolve type=rule control=sequence
//ff:what TestAddPublication — rising versions append; a repeated or lower version fails
package resolve

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestAddPublication(t *testing.T) {
	lg := &versionLog{}
	if addPublication(lg, core.Publication{Version: 1}) != nil || addPublication(lg, core.Publication{Version: 3}) != nil || len(lg.Entries) != 2 {
		t.Fatalf("rising: %+v", lg)
	}
	if addPublication(lg, core.Publication{Version: 3}) == nil {
		t.Fatal("repeat: want error")
	}
}
