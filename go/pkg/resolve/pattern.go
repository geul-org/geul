//ff:type feature=resolve type=model
//ff:what Pattern — `<Issuer>[/<문서 ID>]` or `*`, matching whole segments
package resolve

// Pattern selects documents (GEUL-ORG 문서 패턴). Issuer is a canonical Issuer
// ID or "*" for every Issuer; Doc is a canonical document ID or "" for every
// document of the Issuer. There is no partial-hex wildcard.
type Pattern struct {
	Issuer string
	Doc    string
}
