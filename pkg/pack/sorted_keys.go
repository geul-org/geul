//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what sortedKeys — a key set's keys in ascending order
package pack

import "sort"

// sortedKeys returns the keys of s sorted.
func sortedKeys(s keySet) []string {
	out := make([]string, 0, len(s))
	for k := range s {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
