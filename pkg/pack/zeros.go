//ff:func feature=pack type=rule control=iteration dimension=1
//ff:what zeros — whether every byte of a reserved field is zero
package pack

// zeros reports whether b is all zero (reserved fields must be).
func zeros(b []byte) bool {
	for _, x := range b {
		if x != 0 {
			return false
		}
	}
	return true
}
