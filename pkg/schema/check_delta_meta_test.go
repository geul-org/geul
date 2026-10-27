//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCheckDeltaMeta — well-formed DELTA-BASE, DELTA-REMOVE and NEXT-LOCAL-ID pass; version 0, an empty removal and 0 are M7
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCheckDeltaMeta(t *testing.T) {
	for _, r := range []core.Record{core.EncodeDeltaBase(3), core.EncodeDeltaRemove([]uint64{7}), core.EncodeNextLocalID(9)} {
		if fs := checkDeltaMeta(recOf(r)); fs != nil {
			t.Fatalf("%+v: %+v", r, fs)
		}
	}
	for _, r := range []core.Record{core.EncodeDeltaBase(0), core.EncodeDeltaRemove(nil), core.EncodeNextLocalID(0)} {
		if rulesOf(checkDeltaMeta(recOf(r))) != "M7" {
			t.Fatalf("%+v: want M7", r)
		}
	}
}
