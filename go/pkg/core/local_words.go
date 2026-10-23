//ff:func feature=core type=encoder control=sequence
//ff:what localWords — the words of a LOCAL Reference record
package core

// localWords encodes the LOCAL Reference record of id.
func localWords(id uint64) []Word {
	ws, _ := EncodeRecord(EncodeLocalRef(id))
	return ws
}
