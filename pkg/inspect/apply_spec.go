//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what applySpec — set each entry's codec spec and the language registry
package inspect

import "github.com/geul-org/gogeul/pkg/codecspec"

// applySpec gives each entry the fields s declares for its codec and the
// language registry s carries (Spec.Languages).
func applySpec(entries []Entry, s codecspec.Spec) {
	langs := s.Languages()
	for i := range entries {
		e := &entries[i]
		e.Langs = langs
		fs, ok := s.Of(e.Record)
		if !ok {
			continue
		}
		e.Spec = fs
	}
}
