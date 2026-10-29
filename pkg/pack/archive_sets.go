//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what Archive.Sets — the opened sets in read order with their current generation
package pack

// Sets describes the sets that have a generation.
func (a *Archive) Sets() []SetInfo {
	a.mu.Lock()
	defer a.mu.Unlock()
	var out []SetInfo
	for _, s := range a.sets {
		if s.r != nil {
			out = append(out, SetInfo{Name: s.name, Dir: s.dir, Generation: s.r.m.gen, Manifest: s.r.name, Objects: len(s.r.m.objects)})
		}
	}
	return out
}
