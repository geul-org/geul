//ff:func feature=pack type=resolver control=sequence
//ff:what Reader.hasNamedUnder — whether any NAMED path starts with prefix
package pack

import (
	"sort"
	"strings"
)

// hasNamedUnder reports a named path with the given prefix.
func (r *Reader) hasNamedUnder(prefix string) bool {
	ns := r.m.named
	i := sort.Search(len(ns), func(i int) bool { return ns[i].path >= prefix })
	return i < len(ns) && strings.HasPrefix(ns[i].path, prefix)
}
