//ff:func feature=resolve type=resolver control=sequence
//ff:what nextLocalID — the version list's next LOCAL ID after publishing a document: above every LOCAL it uses, never below the previous value
package resolve

import "github.com/geul-org/gogeul/pkg/inspect"

// nextLocalID is the NEXT-LOCAL-ID a version list records after publishing
// document b (meta/lineage/NEXT-LOCAL-ID): one above the largest LOCAL ID b
// uses (inspect.MaxLocalID), but never below prev, the list's earlier value,
// so a deleted top ID is not given out again. It is at least 1.
func nextLocalID(prev uint64, b []byte) uint64 {
	d, _ := inspect.Decode(b)
	return max(prev, inspect.MaxLocalID(d)+1)
}
