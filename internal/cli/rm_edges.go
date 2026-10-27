//ff:func feature=cli type=command control=iteration dimension=1
//ff:what rmEdges — drop the named Edges of a draft, refusing while a remaining Edge refers to one
package cli

import (
	"fmt"
	"slices"
)

// rmEdges removes each Edge `#n` of args from the draft at path, from the
// BODY or the HEAD (draft.edgeAt). Unless
// force, a remaining record referring to a removed ID (referrer) stops it.
func rmEdges(path string, args []string, force bool) error {
	d, err := readDraft(path)
	if err != nil {
		return err
	}
	var gone []uint64
	for _, s := range args {
		id, err := parseID(s)
		if err != nil {
			return err
		}
		part, i := d.edgeAt(id)
		if i < 0 {
			return fmt.Errorf("no Edge #%d in %s", id, path)
		}
		*part = slices.Delete(*part, i, i+1)
		gone = append(gone, id)
	}
	if by, ok := referrer(d, gone); ok && !force {
		return fmt.Errorf("#%d still refers to a removed Edge (use --force)", by)
	}
	return saveDraft(path, d)
}
