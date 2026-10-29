//ff:type feature=pack type=model
//ff:what streamObject — a reserved object of a Stream: kind, size and hash once closed
package pack

// streamObject is a reserved object (Size 0 while open).
type streamObject struct {
	Kind byte
	Size uint64
	Hash [32]byte
}
