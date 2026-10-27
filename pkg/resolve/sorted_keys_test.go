//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what sortedKeys — test helper: map keys in order
package resolve

import (
	"sort"

	"github.com/geul-org/gogeul/pkg/core"
)

// sortedKeys returns the keys of m sorted, for deterministic fixtures.
func sortedKeys(m map[string]core.Word) []string {
	ks := make([]string, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	return ks
}
