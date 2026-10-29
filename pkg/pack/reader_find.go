//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what Reader.find — key's record location in the newest layer that has it; a deletion marker hides older layers
package pack

import "context"

// find searches layers newest first. The first layer holding key decides: a
// deletion marker is NotFound without looking further.
func (r *Reader) find(ctx context.Context, ls []layer, key, path string) (Loc, error) {
	for _, l := range ls {
		loc, ok, err := r.findInLayer(ctx, l, key)
		if err != nil {
			return Loc{}, err
		}
		if ok {
			return liveLoc(loc, path)
		}
	}
	return Loc{}, notFound(path)
}
