//ff:func feature=page type=helper control=sequence
//ff:what funcField — a FuncRef as its VerbCode plus a link to its ALIAS (none for a local function)
package page

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/toulmin"
)

// funcField shows a FuncRef: its VerbCode and a link to the package ALIAS, or
// "local" when the ALIAS is 0.
func funcField(key string, fn toulmin.FuncRef, labels map[uint64]string) Field {
	f := Field{Key: key, Text: fmt.Sprintf("verb 0x%04X", fn.Verb)}
	if fn.Alias == 0 {
		f.Text += " local"
		return f
	}
	f.Refs = refs([]toulmin.TID{fn.Alias}, labels)
	return f
}
