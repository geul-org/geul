//ff:func feature=core type=decoder control=selection
//ff:what HeadEdge — whether an Edge belongs in the HEAD: a SOURCE on the whole document (Target LOCAL 0) or a QUALIFIER on such a SOURCE; with its Edge Id
package core

// HeadEdge reports whether r is an Edge the HEAD holds (document HEAD):
// a direct-code SOURCE whose Target is LOCAL 0, the source of the whole
// document (edge/SOURCE), or a direct-code QUALIFIER whose Target is in
// sources, the Edge Ids of the HEAD SOURCEs before it (edge/QUALIFIER). It
// returns r's Edge Id. Every other record, Meta included, is false.
func HeadEdge(r Record, sources map[uint64]bool) (uint64, bool) {
	id, ok := localAt(r.Payload, 0)
	switch {
	case r.Kind != KindEdge || len(r.Issuer) > 0 || !ok:
		return 0, false
	case r.Codec == EdgeSource:
		t, ok := localAt(r.Payload, 1)
		return id, ok && t == 0
	case r.Codec == EdgeQualifier:
		t, ok := localAt(r.Payload, 2)
		return id, ok && sources[t]
	default:
		return 0, false
	}
}
