//ff:func feature=inspect type=encoder control=sequence
//ff:what tombstoneSpec — test fixture: TOMBSTONE declared as doc 4, version 2, reason 1 (code 1 = license)
package inspect

import (
	"github.com/geul-org/geul/go/pkg/codecspec"
	"github.com/geul-org/geul/go/pkg/core"
)

// tombstoneSpec declares TOMBSTONE for layout tests.
func tombstoneSpec() codecspec.Spec {
	return codecspec.Spec{core.TargetWord(core.KindEdge, core.EdgeTombstone): {
		{Name: "doc", Width: 4}, {Name: "version", Width: 2}, {Name: "reason", Width: 1, Enums: map[string]string{"0x0001": "license"}},
	}}
}
