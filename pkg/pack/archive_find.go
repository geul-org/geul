//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what Archive.find — the set named name, nil when there is none
package pack

// find looks a set up by name.
func (a *Archive) find(name string) *packSet {
	for _, s := range a.sets {
		if s.name == name {
			return s
		}
	}
	return nil
}
