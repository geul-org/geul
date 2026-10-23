//ff:func feature=core type=decoder control=sequence
//ff:what takeLocal — the LOCAL Reference record at the start of ws: its ID and word count
package core

// takeLocal decodes the LOCAL Reference record that starts ws.
func takeLocal(ws []Word) (uint64, int, error) {
	r, n, err := DecodeRecord(ws)
	if err != nil {
		return 0, 0, err
	}
	id, err := DecodeLocalRef(r)
	return id, n, err
}
