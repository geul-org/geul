//ff:type feature=pack type=model
//ff:what object — one OBJECTS entry: kind (data pack or index object), size and SHA-256 (= its name)
package pack

// object is an immutable stored object named by the SHA-256 of its bytes.
type object struct {
	kind byte
	size uint64
	hash [32]byte
}
