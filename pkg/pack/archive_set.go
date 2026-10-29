//ff:func feature=pack type=writer control=sequence
//ff:what Archive.set — the set named name, created (empty, kept in read order) when missing
package pack

import (
	"path/filepath"
	"sort"
)

// set returns the set for name ("" is the top set), adding it if needed.
func (a *Archive) set(name string) *packSet {
	if s := a.find(name); s != nil {
		return s
	}
	s := &packSet{name: name, dir: filepath.Join(a.dir, name)}
	a.sets = append(a.sets, s)
	sort.SliceStable(a.sets, func(i, j int) bool {
		return a.sets[i].name == "" && a.sets[j].name != "" || a.sets[j].name != "" && issuerLess(a.sets[i].name, a.sets[j].name)
	})
	return s
}
