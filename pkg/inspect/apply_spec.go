//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what applySpec — set each entry's codec spec
package inspect

import "github.com/geul-org/gogeul/pkg/codecspec"

// applySpec gives each entry the fields s declares for its codec.
func applySpec(entries []Entry, s codecspec.Spec) {
	for i := range entries {
		e := &entries[i]
		fs, ok := s.Of(e.Record)
		if !ok {
			continue
		}
		e.Spec = fs
	}
}
