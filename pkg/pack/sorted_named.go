//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what sortedNamed — NAMED entries in path byte order; a path put twice is an error
package pack

import (
	"fmt"
	"sort"
)

// sortedNamed sorts the named records by path (§4.3 NAMED: ascending, no
// duplicates).
func sortedNamed(kvs []streamKV) ([]namedEntry, error) {
	out := make([]namedEntry, len(kvs))
	for i, e := range kvs {
		out[i] = namedEntry{path: e.Key, loc: e.Loc}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].path < out[j].path })
	for i := 1; i < len(out); i++ {
		if out[i].path == out[i-1].path {
			return nil, fmt.Errorf("geul/pack: named path %s put twice", out[i].path)
		}
	}
	return out, nil
}
