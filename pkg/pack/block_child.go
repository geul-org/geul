//ff:func feature=pack type=resolver control=sequence
//ff:what block.child — the last child of an internal block whose first key ≤ key, -1 when none
package pack

import "sort"

// child returns the index of the child covering key: the last first-key ≤ key.
func (b block) child(key string) int {
	return sort.Search(len(b.keys), func(i int) bool { return b.keys[i] > key }) - 1
}
