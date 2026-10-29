//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what Archive.Commit — write the staged changes as one new generation per changed set, reopen those sets, clear the stage
package pack

import (
	"context"
	"sort"
)

// Commit records the staged changes. Sets are written in read order; a set
// whose bytes did not change gets no generation. On error the stage is kept.
func (a *Archive) Commit(ctx context.Context) ([]BuildResult, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if len(a.puts)+len(a.dels) == 0 {
		return nil, nil
	}
	bs, err := a.batches(ctx)
	if err != nil {
		return nil, err
	}
	sets := make([]*packSet, 0, len(bs))
	for s := range bs {
		sets = append(sets, s)
	}
	sort.Slice(sets, func(i, j int) bool { return sets[i].name < sets[j].name })
	var out []BuildResult
	for _, s := range sets {
		res, err := a.commitSet(ctx, s, bs[s])
		if err != nil {
			return out, err
		}
		out = append(out, res)
	}
	a.puts, a.dels = map[string][]byte{}, map[string]bool{}
	return out, nil
}
