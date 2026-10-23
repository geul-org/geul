//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what edgeIDDups — E2: no Edge Id is declared twice
package schema

import "fmt"

// edgeIDDups reports every Edge Id declared by more than one Edge.
func edgeIDDups(body []rec) []Fact {
	seen := map[uint64]bool{}
	var fs []Fact
	for _, r := range body {
		id, ok := edgeID(r)
		if !ok {
			continue
		}
		if seen[id] {
			fs = append(fs, factAt("E2", Error, r, "Edge Id", "문서 안에서 유일", fmt.Sprintf("%d 중복", id)))
		}
		seen[id] = true
	}
	return fs
}
