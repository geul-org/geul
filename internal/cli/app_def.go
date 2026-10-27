//ff:func feature=cli type=command control=sequence
//ff:what app.def — append REFERENCE-NODE(#edge, #node, object) with two new IDs and print the node
package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/geul-org/gogeul/pkg/notation"
)

// def takes the next two free IDs, the Edge's and the node's, and appends a
// REFERENCE-NODE binding the node to the object tokens.
func (a *app) def(path string, object []string) error {
	d, err := readDraft(path)
	if err != nil {
		return err
	}
	next := a.nextFree(d)
	e, err := notation.ParseEdge("REFERENCE-NODE #"+strconv.FormatUint(next+1, 10)+" "+strings.Join(object, " "), a.context())
	if err != nil {
		return err
	}
	d.body = append(d.body, edgeRecord(next, e))
	if err := saveDraft(path, d); err != nil {
		return err
	}
	fmt.Fprintf(a.out, "#%d\n", next+1)
	return nil
}
