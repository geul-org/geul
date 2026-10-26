//ff:func feature=cli type=command control=sequence
//ff:what app.hash — print a document's content hash as 0x and 64 hex digits
package cli

import (
	"encoding/hex"
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// hash prints core.ContentHash of the document (readDoc): what its version
// list records and its `h/` file name.
func (a *app) hash(arg string) error {
	b, err := a.readDoc(arg)
	if err != nil {
		return err
	}
	ws, err := core.ReadWords(b)
	if err != nil {
		return err
	}
	h, err := core.ContentHash(ws)
	if err != nil {
		return err
	}
	fmt.Fprintln(a.out, "0x"+hex.EncodeToString(h[:]))
	return nil
}
