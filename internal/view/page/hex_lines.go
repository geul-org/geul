//ff:func feature=page type=helper control=iteration dimension=1
//ff:what hexLines — words as one line each, the first named; a long run cut after hexLimit words
package page

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// hexLimit is the number of words a field shows before the rest is cut.
const hexLimit = 32

// hexLines renders ws one word per line, naming the first line name. More than
// hexLimit words show the first hexLimit and a "…" line saying how many follow.
func hexLines(ws []core.Word, name string) []HexLine {
	out := make([]HexLine, 0, min(len(ws), hexLimit)+1)
	for i, w := range ws {
		if i == hexLimit {
			return append(out, HexLine{Word: "…", Note: fmt.Sprintf("+%d WORDs", len(ws)-hexLimit)})
		}
		out = append(out, HexLine{Word: fmt.Sprintf("%04X", uint16(w)), Name: name})
		name = ""
	}
	return out
}
