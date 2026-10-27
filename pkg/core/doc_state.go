//ff:type feature=core type=model
//ff:what docState — the HEAD/BODY/END position while SplitDocument walks a document's records
package core

// docState tracks where SplitDocument is inside a document: still in the HEAD,
// in the BODY (after BODY, or immediately when the HEAD is empty and BODY is
// omitted), or closed by END.
type docState struct {
	inBody bool
	ended  bool
	doc    Document
}
