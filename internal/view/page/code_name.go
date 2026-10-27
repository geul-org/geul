//ff:func feature=page type=helper control=sequence
//ff:what codeName — a header code's name, or its number when unnamed
package page

import "strconv"

// codeName returns the name of code c in names, or c in decimal.
func codeName[C ~uint16](names map[C]string, c C) string {
	if n, ok := names[c]; ok {
		return n
	}
	return strconv.FormatUint(uint64(c), 10)
}
