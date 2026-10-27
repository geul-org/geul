//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what payloadParts — a record payload as Parts: the typed layout at the SPAN width that fits, else the generic split
package inspect

// payloadParts lays out e's payload by its codec spec when one fits, else from
// its typed content. The SPAN width is
// not stored on the entry, so each declarable width is tried and the first
// layout that uses every word exactly is kept. Content with no typed layout,
// or no fitting width, falls back to genericParts.
func payloadParts(e Entry) []Part {
	if ps, ok := specParts(e); ok {
		return ps
	}
	for _, sw := range []int{1, 2, 4, 8} {
		w := &walker{ws: e.Record.Payload, sw: sw}
		if typedLayout(w, e.Data) && !w.bad && w.pos == len(w.ws) {
			return w.out
		}
	}
	return genericParts(e)
}
