//ff:func feature=geulmd type=encoder control=sequence
//ff:what funcStr — a FuncRef as `#alias/0xverb`
package geulmd

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/toulmin"
)

// funcStr renders a FuncRef as its ALIAS reference and VerbCode: `#alias/0xverb`
// (`#0` is a local function).
func funcStr(f toulmin.FuncRef) string {
	return fmt.Sprintf("#%d/0x%04x", f.Alias, f.Verb)
}
