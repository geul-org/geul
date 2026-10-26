//ff:func feature=resolve type=helper control=sequence
//ff:what hexWords — the words of a canonical hex ID already checked by parseHexID
package resolve

import "github.com/geul-org/geul/go/pkg/core"

// hexWords turns the bytes parseHexID decoded into words.
func hexWords(b []byte) []core.Word {
	ws, _ := core.ReadWords(b)
	return ws
}
