//ff:type feature=resolve type=model
//ff:what Named — an Issuer or document ID found in an archive with the names bound to it, main name last
package resolve

// Named is an ID in canonical hex and its names in binding order; the last is
// the main name (edge/PUBLICATION 문서 이름).
type Named struct {
	ID    string   `json:"id"`
	Names []string `json:"names,omitempty"`
}
