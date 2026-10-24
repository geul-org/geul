//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what ParseIssuerNames — the name → Issuer map of an issuers.geul document (BIND records)
package resolve

import "github.com/geul-org/geul/go/pkg/core"

// ParseIssuerNames reads issuers.geul (GEUL-ORG Issuer 이름): every BIND whose
// Subject is ISSUER-POINTER(Issuer) maps its name to that Issuer in
// canonical hex. Other records are ignored. A later BIND of the same name wins.
func ParseIssuerNames(b []byte) (map[string]string, error) {
	ws, err := core.ReadWords(b)
	if err != nil {
		return nil, err
	}
	doc, err := core.SplitDocument(ws)
	if err != nil {
		return nil, err
	}
	names := map[string]string{}
	for _, r := range doc.Body {
		bind, err := core.DecodeBind(r)
		if err != nil {
			continue
		}
		if issuer, ok := issuerOf(bind.Subject); ok {
			names[string(bind.Name)] = issuer
		}
	}
	return names, nil
}
