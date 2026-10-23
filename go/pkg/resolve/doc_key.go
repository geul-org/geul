//ff:type feature=resolve type=model
//ff:what DocKey — a document's identity apart from its version, as canonical hex IDs
package resolve

// DocKey is a document's identity regardless of version: the Issuer and the
// document ID in canonical hex (GEUL-ORG ID 표기). A graph holds one Node per DocKey.
type DocKey struct {
	Issuer string
	Doc    string
}
