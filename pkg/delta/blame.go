//ff:func feature=delta type=decoder control=iteration dimension=1
//ff:what Blame — for each Edge of the last version, the version that added its Edge Id and the one that last changed it
package delta

import "errors"

// Blame walks vs (oldest first, snapshots of one document) and returns the
// last version's Edges in order, each with the version that first held its
// Edge Id and the last version that changed its content.
func Blame(vs []Version) ([]Line, error) {
	if len(vs) == 0 {
		return nil, errors.New("geul/delta: no versions to blame")
	}
	seen := map[uint64]*Line{}
	var last snapshot
	for _, v := range vs {
		s, err := splitSnapshot(v.Doc)
		if err != nil {
			return nil, err
		}
		blameStep(seen, s, v.V)
		last = s
	}
	out := make([]Line, len(last.ids))
	for i, id := range last.ids {
		out[i] = *seen[id]
	}
	return out, nil
}
