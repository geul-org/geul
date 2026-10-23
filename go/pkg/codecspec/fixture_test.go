//ff:func feature=codecspec type=encoder control=sequence
//ff:what sampleBody — test fixture: TOMBSTONE declared with named doc, version and reason fields and one reason code
package codecspec

import (
	"github.com/geul-org/geul/go/pkg/core"
)

// sampleBody declares TOMBSTONE: doc (4), version (2), reason (1, code 1 = "license").
func sampleBody() []core.Record {
	return []core.Record{
		core.EncodeFieldDecl(core.FieldDecl{EdgeID: 2, Width: 4}),
		core.EncodeFieldDecl(core.FieldDecl{EdgeID: 3, Width: 2}),
		core.EncodeFieldDecl(core.FieldDecl{EdgeID: 4, Width: 1}),
		core.EncodeEnumDecl(core.EnumDecl{EdgeID: 5, Field: 4, Code: []core.Word{1}}),
		core.EncodeDeclare(core.Declare{EdgeID: 1, Target: core.TargetWord(core.KindEdge, core.EdgeTombstone), Fields: []uint64{2, 3, 4}}),
		nameOf(6, 1, "TOMBSTONE"), nameOf(7, 2, "doc"), nameOf(8, 3, "version"), nameOf(9, 4, "reason"), nameOf(10, 5, "license"),
	}
}
