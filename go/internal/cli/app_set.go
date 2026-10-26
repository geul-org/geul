//ff:func feature=cli type=command control=sequence
//ff:what app.set — read new fields for an Edge under its own codec name and replace it in place
package cli

import (
	"fmt"
	"strings"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/notation"
)

// set replaces Edge idArg's payload after its Edge Id with fields read as
// a line of its codec (notation.ParseEdge), keeping its place and ID.
func (a *app) set(path, idArg string, fields []string) error {
	id, err := parseID(idArg)
	if err != nil {
		return err
	}
	d, err := readDraft(path)
	if err != nil {
		return err
	}
	i := edgeIndex(d.body, id)
	if i < 0 {
		return fmt.Errorf("no Edge #%d in %s", id, path)
	}
	name := core.CodecName(core.KindEdge, d.body[i].Codec)
	e, err := notation.ParseEdge(name+" "+strings.Join(fields, " "), a.context())
	if err != nil {
		return err
	}
	d.body[i] = edgeRecord(id, e)
	return saveDraft(path, d)
}
