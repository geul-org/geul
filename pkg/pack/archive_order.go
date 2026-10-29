//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what Archive.order — the sets to search for a path: its Issuer's set then the top set for `<Issuer>/…`, every set otherwise
package pack

import "strings"

// order lists the sets that may serve name, most specific first.
func (a *Archive) order(name string) []*packSet {
	first, _, _ := strings.Cut(name, "/")
	var own, rest []*packSet
	for _, s := range a.sets {
		switch {
		case s.r == nil:
		case s.name != "" && s.name == first:
			own = append(own, s)
		case s.name == "" || !isIssuer(first):
			rest = append(rest, s)
		}
	}
	return append(own, rest...)
}
