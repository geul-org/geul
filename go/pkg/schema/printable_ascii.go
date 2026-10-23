//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what printableASCII — every byte is visible ASCII (0x21~0x7E)
package schema

// printableASCII reports whether every byte of b is in 0x21~0x7E.
func printableASCII(b []byte) bool {
	for _, c := range b {
		if c < 0x21 || c > 0x7E {
			return false
		}
	}
	return true
}
