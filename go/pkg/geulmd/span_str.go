//ff:func feature=geulmd type=encoder control=sequence
//ff:what spanStr — a source span as `@offset+length`
package geulmd

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/toulmin"
)

// spanStr renders a SPAN as `@offset+length` in bytes of the SOURCE text.
func spanStr(s toulmin.Span) string {
	return fmt.Sprintf("@%d+%d", s.Off, s.Len)
}
