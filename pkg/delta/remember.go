//ff:func feature=delta type=decoder control=sequence
//ff:what LocalIDs.remember — note one previous BODY record: an Edge by its content key, a REFERENCE-NODE also by its Object
package delta

import "github.com/geul-org/gogeul/pkg/core"

// remember notes r, a BODY record of the previous version. A record that is
// not an Edge with a LOCAL Edge Id is ignored (it has no ID to carry).
func (l *LocalIDs) remember(r core.Record) {
	id, rest, ok := edgeID(r)
	if !ok {
		return
	}
	l.prev = append(l.prev, id)
	l.edges[edgeKey(r.Codec, rest)] = append(l.edges[edgeKey(r.Codec, rest)], id)
	if r.Codec != edgeReferenceNode {
		return
	}
	subject, n, err := core.DecodeRecord(rest)
	if local, lerr := core.DecodeLocalRef(subject); err == nil && lerr == nil {
		l.defs[edgeKey(r.Codec, rest[n:])] = [2]uint64{id, local}
	}
}
