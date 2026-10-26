//ff:type feature=resolve type=model
//ff:what History — a document's published versions in order and its NEXT-LOCAL-ID
package resolve

// History is what a document's version list says: every published version in
// order with its state and time, and the NEXT-LOCAL-ID (0 when absent).
type History struct {
	Entries []LogEntry
	Next    uint64
}
