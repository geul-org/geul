//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what children — the direct entries of a directory in a sorted path list, skipping over each subdirectory's contents
package pack

import (
	"sort"
	"strings"
)

// children lists the entries under prefix ("" for the top, else "dir/") in
// sorted paths.
func children(paths []string, prefix string) []dirEntry {
	var out []dirEntry
	i := sort.SearchStrings(paths, prefix)
	for i < len(paths) && strings.HasPrefix(paths[i], prefix) {
		seg, _, more := strings.Cut(paths[i][len(prefix):], "/")
		out = append(out, dirEntry{name: seg, dir: more})
		i++
		if more {
			i = sort.SearchStrings(paths, prefix+seg+"0") // '0' follows '/': past "seg/…"
		}
	}
	return out
}
