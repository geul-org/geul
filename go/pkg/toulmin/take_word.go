//ff:func feature=toulmin type=decoder control=sequence
//ff:what TakeWord — read one code word at *pos and advance past it
package toulmin

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// TakeWord reads the single word at ws[*pos] (a header, code or count word) and
// advances *pos. what names the field for the truncation error.
func TakeWord(ws []core.Word, pos *int, what string) (core.Word, error) {
	if *pos >= len(ws) {
		return 0, fmt.Errorf("geul/toulmin: %s word missing", what)
	}
	w := ws[*pos]
	*pos++
	return w, nil
}
