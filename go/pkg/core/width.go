//ff:func feature=core type=helper control=selection
//ff:what WidthWords — the word count of a 2-bit width code (00=1, 01=2, 10=4, 11=8 WORD)
package core

// WidthWords returns the word count of a 2-bit width code as used by the Meta
// ID-WIDTH fields and the TANGEUL SPAN width: 00=1, 01=2, 10=4, 11=8 WORD.
func WidthWords(code uint8) int {
	switch code & 3 {
	case 0:
		return 1
	case 1:
		return 2
	case 2:
		return 4
	default:
		return 8
	}
}
