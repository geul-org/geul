//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what writer.planDeletes — deletions the previous generation needs: listed ones, or all unlisted ones for a Full batch
package pack

import "sort"

// planDeletes fills w.dels with previous paths to remove, sorted.
func (w *writer) planDeletes() {
	gone := map[string]bool{}
	for _, d := range w.b.Delete {
		gone[d] = true
	}
	for p := range w.base {
		_, kept := w.sums[p]
		if !kept && (w.b.Full || gone[p]) {
			w.dels = append(w.dels, p)
		}
	}
	sort.Strings(w.dels)
}
