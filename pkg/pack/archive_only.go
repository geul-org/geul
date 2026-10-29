//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what Archive.only — the single Issuer set that staged `<Issuer>/…` paths go to, nil when none or several
package pack

import (
	"context"
	"strings"
)

// only routes the staged Issuer paths and returns their set when it is one.
func (a *Archive) only(ctx context.Context, names []string) (*packSet, error) {
	var one *packSet
	for _, n := range names {
		if first, _, nested := strings.Cut(n, "/"); !isIssuer(first) || !nested {
			continue
		}
		s, err := a.target(ctx, n, nil)
		if err != nil || one != nil && s != one {
			return nil, err
		}
		one = s
	}
	return one, nil
}
