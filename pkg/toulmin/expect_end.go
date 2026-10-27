//ff:func feature=toulmin type=decoder control=sequence
//ff:what ExpectEnd — require that a record's fields consumed its whole payload
package toulmin

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// ExpectEnd reports an error when the fields read so far (pos words) do not
// account for the whole payload — trailing words mean the layout and the
// record's Length disagree.
func ExpectEnd(ws []core.Word, pos int, what string) error {
	if pos != len(ws) {
		return fmt.Errorf("geul/toulmin: %s fields use %d of %d payload words", what, pos, len(ws))
	}
	return nil
}
