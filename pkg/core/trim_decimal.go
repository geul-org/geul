//ff:func feature=core type=encoder control=iteration dimension=1
//ff:what TrimDecimal — strip trailing zero digits from C×10^E (0 becomes C=0, E=0)
package core

// TrimDecimal removes trailing decimal zeros from the coefficient, raising the
// exponent to match, so 8000×10^-4 becomes 8×10^-1 (GEUL t002 TOULMIN: the
// encoder writes the shortest coefficient). Zero normalizes to 0×10^0.
func TrimDecimal(e int32, c int64) (int32, int64) {
	if c == 0 {
		return 0, 0
	}
	for c%10 == 0 {
		c /= 10
		e++
	}
	return e, c
}
