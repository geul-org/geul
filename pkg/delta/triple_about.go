//ff:func feature=delta type=decoder control=sequence
//ff:what tripleAbout — a TRIPLE whose Subject is a LOCAL node: that LOCAL and the TRIPLE's words without Edge Id and Subject
package delta

import (
	"slices"

	"github.com/geul-org/gogeul/pkg/core"
)

// tripleAbout splits r when it is a TRIPLE whose Subject is a LOCAL Reference
// other than 0: it returns the Subject's ID and the Predicate words followed
// by the Object and Rank words (what the TRIPLE says of that LOCAL).
func tripleAbout(r core.Record) (uint64, []core.Word, bool) {
	_, rest, ok := edgeID(r)
	if !ok || r.Codec != edgeTriple || len(r.Issuer) > 0 {
		return 0, nil, false
	}
	_, p, err := core.DecodeRecord(rest)
	if err != nil {
		return 0, nil, false
	}
	subject, s, err := core.DecodeRecord(rest[p:])
	id, lerr := core.DecodeLocalRef(subject)
	if err != nil || lerr != nil || id == 0 {
		return 0, nil, false
	}
	return id, slices.Concat(rest[:p], rest[p+s:]), true
}
