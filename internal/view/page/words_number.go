//ff:func feature=page type=helper control=iteration dimension=1
//ff:what wordsNumber — payload words read as one big-endian unsigned number
package page

import "github.com/geul-org/gogeul/pkg/core"

// wordsNumber reads ws as one big-endian number, most significant word first.
// The caller keeps ws to at most 4 words.
func wordsNumber(ws []core.Word) uint64 {
	var n uint64
	for _, w := range ws {
		n = n<<16 | uint64(w)
	}
	return n
}
