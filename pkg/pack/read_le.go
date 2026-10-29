//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what readLE — an unsigned little-endian integer of up to 8 bytes
package pack

// readLE reads b as little-endian.
func readLE(b []byte) uint64 {
	var v uint64
	for i := len(b) - 1; i >= 0; i-- {
		v = v<<8 | uint64(b[i])
	}
	return v
}
