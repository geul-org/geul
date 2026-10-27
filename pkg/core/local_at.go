//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what localAt — the ID of the LOCAL Reference that is record i of a payload; false when the records before it or it do not decode, or it is not a LOCAL
package core

// localAt walks ws record by record and decodes record i (0 is the first)
// as a LOCAL Reference (DecodeLocalRef).
func localAt(ws []Word, i int) (uint64, bool) {
	for ; i > 0; i-- {
		_, n, err := DecodeRecord(ws)
		if err != nil {
			return 0, false
		}
		ws = ws[n:]
	}
	id, _, err := takeLocal(ws)
	return id, err == nil
}
