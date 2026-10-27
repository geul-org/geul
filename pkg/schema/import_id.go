//ff:func feature=schema type=decoder control=iteration dimension=1
//ff:what importID — the IMPORT ID of an IMPORT-DOC, when it is well formed and fits 64 bits
package schema

// importID reads the IMPORT ID of an IMPORT-DOC record: the words after its
// DOC-POINTER.
func importID(r rec) (uint64, bool) {
	ws, ok := importSplit(r)
	if !ok || len(ws) < 1 || len(ws) > 4 {
		return 0, false
	}
	var v uint64
	for _, w := range ws {
		v = v<<16 | uint64(w)
	}
	return v, true
}
