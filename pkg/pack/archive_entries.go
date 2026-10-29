//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what Archive.entries — the listing of a directory with each file entry tied to the Archive for Info
package pack

// entries lists the children of the fs directory name in sorted paths idx.
func (a *Archive) entries(idx []string, name string) []dirEntry {
	es := children(idx, dirPrefix(name))
	for i := range es {
		es[i].a, es[i].path = a, dirPrefix(name)+es[i].name
	}
	return es
}
