//ff:type feature=core type=model
//ff:what docState — the HEAD/BODY/END position while SplitDocument walks a document's records, with the HEAD SOURCE Edge Ids
package core

// docState tracks where SplitDocument is inside a document: still in the HEAD,
// in the BODY (after BODY, or immediately when the HEAD is empty and BODY is
// omitted), or closed by END. sources holds the Edge Ids of the SOURCEs the
// HEAD has so far (placeHeadEdge): a QUALIFIER on one of them is a HEAD
// record.
type docState struct {
	inBody  bool
	ended   bool
	sources map[uint64]bool
	doc     Document
}
