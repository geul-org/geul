//ff:func feature=pack type=resolver control=sequence
//ff:what liveLoc — a found location, or NotFound when it is a deletion marker
package pack

// liveLoc turns a deletion marker into notFound(path).
func liveLoc(l Loc, path string) (Loc, error) {
	if l.Enc == encGone {
		return Loc{}, notFound(path)
	}
	return l, nil
}
