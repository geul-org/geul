//ff:func feature=schema type=helper control=sequence
//ff:what fmtWords — "N WORD"
package schema

import "strconv"

// fmtWords renders a word count.
func fmtWords(n int) string {
	return strconv.Itoa(n) + " WORD"
}
