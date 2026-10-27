//ff:func feature=resolve type=resolver control=selection
//ff:what Dir.issuerOfName — an Issuer written as a canonical ID or as a name of issuers.geul
package resolve

import (
	"context"
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// issuerOfName returns the Issuer s names: a canonical ID as is, otherwise
// the name looked up in `issuers.geul`.
func (d Dir) issuerOfName(s string) ([]core.Word, error) {
	if isIssuerID(s) {
		b, _ := parseHexID(s, (len(s)-2)/4)
		return hexWords(b), nil
	}
	b, err := d.Fetch(context.Background(), issuersFile)
	if err != nil {
		return nil, fmt.Errorf("geul/resolve: Issuer name %q: %w", s, err)
	}
	names, err := ParseIssuerNames(b)
	if err != nil {
		return nil, err
	}
	id, ok := names[s]
	if !ok {
		return nil, fmt.Errorf("%w: Issuer name %q", ErrNotFound, s)
	}
	h, _ := parseHexID(id, (len(id)-2)/4)
	return hexWords(h), nil
}
