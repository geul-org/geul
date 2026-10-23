//ff:type feature=core type=model
//ff:what Bind — a BIND Edge: Edge Id, Subject Reference and the bound name
package core

// Bind is a BIND Edge (GEUL t002 docs/edge/BIND): the Subject Reference record
// and the machine-matched name bound to it. Name is the UTF-8 bytes as written.
type Bind struct {
	EdgeID  uint64
	Subject Record
	Name    []byte
}
