//ff:func feature=cli type=command control=sequence
//ff:what app.set — read new fields for an Edge under its own codec name and replace it in place
package cli

import (
	"fmt"
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/notation"
)

// set replaces Edge idArg's payload after its Edge Id with fields read as
// a line of its codec (notation.ParseEdge), keeping its place (BODY or
// HEAD, draft.edgeAt) and ID.
func (a *app) set(path, idArg string, fields []string) error {
	id, err := parseID(idArg)
	if err != nil {
		return err
	}
	d, err := readDraft(path)
	if err != nil {
		return err
	}
	part, i := d.edgeAt(id)
	if i < 0 {
		return fmt.Errorf("no Edge #%d in %s", id, path)
	}
	name := core.CodecName(core.KindEdge, (*part)[i].Codec)
	e, err := notation.ParseEdge(name+" "+strings.Join(fields, " "), a.context())
	if err != nil {
		return err
	}
	(*part)[i] = edgeRecord(id, e)
	return saveDraft(path, d)
}
