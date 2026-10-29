//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what OpenDir — the writable Archive of the pack sets under a `.well-known/geul/` directory (none yet is an empty archive)
package pack

import "context"

// OpenDir opens every set under wk at its latest generation. Writes go
// through o (BuildOptions) on Commit.
func OpenDir(ctx context.Context, wk string, o BuildOptions) (*Archive, error) {
	sets, err := findSets(wk)
	if err != nil {
		return nil, err
	}
	for _, s := range sets {
		if s.r, err = Open(ctx, FileStore{Root: s.dir}, Options{}); err != nil {
			return nil, err
		}
	}
	return &Archive{dir: wk, o: o, sets: sets, puts: map[string][]byte{}, dels: map[string]bool{}}, nil
}
