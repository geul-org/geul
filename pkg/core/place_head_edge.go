//ff:func feature=core type=decoder control=sequence
//ff:what placeHeadEdge — put a HEAD Edge (SOURCE on the whole document or its QUALIFIER) into the HEAD, noting a SOURCE's Edge Id; false for any other record
package core

// placeHeadEdge appends r to st's HEAD when it is a HEAD Edge (HeadEdge)
// and, when r is a SOURCE, notes its Edge Id so the QUALIFIERs after it on
// that SOURCE also stay in the HEAD.
func placeHeadEdge(st *docState, r Record) bool {
	id, ok := HeadEdge(r, st.sources)
	if !ok {
		return false
	}
	if r.Codec == EdgeSource {
		st.sources = addID(st.sources, id)
	}
	st.doc.Head = append(st.doc.Head, r)
	return true
}
