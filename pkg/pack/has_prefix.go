//ff:func feature=pack type=helper control=sequence
//ff:what hasPrefix — whether a sorted path list has any path starting with prefix
package pack

import (
	"sort"
	"strings"
)

// hasPrefix binary-searches sorted paths for prefix.
func hasPrefix(paths []string, prefix string) bool {
	i := sort.SearchStrings(paths, prefix)
	return i < len(paths) && strings.HasPrefix(paths[i], prefix)
}
