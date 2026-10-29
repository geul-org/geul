//ff:type feature=pack type=model
//ff:what objectHeader — the 64-byte head of a `.gpk` or `.gpx`: kind, size, inventory place, record or block count
package pack

// objectHeader is §4.4.
type objectHeader struct {
	kind   byte
	size   uint64
	invOff uint64
	invLen uint64
	count  uint64
}
