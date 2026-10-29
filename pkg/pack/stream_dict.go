//ff:type feature=pack type=model
//ff:what streamDict — the dictionary a Stream's records use: ID (0: none), location, hash
package pack

// streamDict is the dictionary the stream's records use (ID 0: none).
type streamDict struct {
	ID   uint32
	Loc  Loc
	Hash [32]byte
}
