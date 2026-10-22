//ff:func feature=core type=helper control=selection
//ff:what WidthCode — the 2-bit width code of a 1/2/4/8-word field width
package core

import "fmt"

// WidthCode is the inverse of WidthWords: 1→00, 2→01, 4→10, 8→11. Any other
// width is an error.
func WidthCode(words int) (uint8, error) {
	switch words {
	case 1:
		return 0, nil
	case 2:
		return 1, nil
	case 4:
		return 2, nil
	case 8:
		return 3, nil
	default:
		return 0, fmt.Errorf("geul/core: width %d is not 1, 2, 4 or 8 words", words)
	}
}
