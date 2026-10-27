//ff:func feature=inspect type=helper control=iteration dimension=1
//ff:what maxOf — the largest of some unsigned numbers
package inspect

// maxOf returns the largest of first and rest.
func maxOf(first uint64, rest ...uint64) uint64 {
	for _, n := range rest {
		if n > first {
			first = n
		}
	}
	return first
}
