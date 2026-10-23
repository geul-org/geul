//ff:func feature=resolve type=decoder control=sequence
//ff:what parsePattern — one `<Issuer>[/<문서 ID>]` or `*` pattern in canonical hex
package resolve

import (
	"fmt"
	"strings"
)

// parsePattern reads one pattern. Each segment is a canonical ID or `*`.
func parsePattern(s string) (Pattern, error) {
	issuer, doc, _ := strings.Cut(s, "/")
	if doc == "*" {
		doc = ""
	}
	_, docOK := parseHexID(doc, 4)
	if (issuer != "*" && !isIssuerID(issuer)) || (doc != "" && !docOK) {
		return Pattern{}, fmt.Errorf("geul/resolve: bad pattern %q", s)
	}
	return Pattern{Issuer: issuer, Doc: doc}, nil
}
