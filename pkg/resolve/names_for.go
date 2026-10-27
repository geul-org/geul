//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what namesFor — every name of a name map that points at one ID, sorted
package resolve

import "sort"

// namesFor returns the names in names mapped to id, sorted.
func namesFor(names map[string]string, id string) []string {
	var out []string
	for n, v := range names {
		if v == id {
			out = append(out, n)
		}
	}
	sort.Strings(out)
	return out
}
