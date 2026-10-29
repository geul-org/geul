//ff:type feature=pack type=model
//ff:what dictEntry — a DICTS entry: zstd dictionary ID, where its bytes lie, their SHA-256
package pack

// dictEntry locates a zstd dictionary stored as a record.
type dictEntry struct {
	id   uint32
	loc  Loc
	hash [32]byte
}
