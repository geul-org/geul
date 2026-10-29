//ff:type feature=pack type=model
//ff:what invEntry — one 32-byte inventory entry: kind, encoding, flags, reference, document, slot, length, offset
package pack

// invEntry is §4.5. ref is the pack Issuer number (kind 1·5), the path table
// number (kind 3·6) or the dictionary ID (kind 4).
type invEntry struct {
	kind  byte
	enc   byte
	flags uint16
	ref   uint32
	doc   uint64
	slot  uint32
	n     uint32
	off   uint64
}
