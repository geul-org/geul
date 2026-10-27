//ff:func feature=cli type=command control=iteration dimension=1
//ff:what app.meta — replace a draft's HEAD Meta records of one codec with a new one, or remove them
package cli

import (
	"errors"
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/notation"
)

// meta reads `NAME VALUE…` (notation.ParseMeta). The new record takes the
// place of the first record of that codec (or goes last); the others of that
// codec are dropped. With unset they are all dropped.
func (a *app) meta(path, name string, values []string, unset bool) error {
	r, err := notation.ParseMeta(strings.TrimSpace(name+" "+strings.Join(values, " ")), a.context())
	if err != nil {
		return err
	}
	if !unset && len(values) == 0 {
		return errors.New("meta: give a value, or --unset")
	}
	d, err := readDraft(path)
	if err != nil {
		return err
	}
	var head []core.Record
	placed := unset
	for _, h := range d.head {
		switch {
		case h.Kind != core.KindMeta || h.Codec != r.Codec:
			head = append(head, h)
		case !placed:
			head, placed = append(head, r), true
		}
	}
	if !placed {
		head = append(head, r)
	}
	d.head = head
	return saveDraft(path, d)
}
