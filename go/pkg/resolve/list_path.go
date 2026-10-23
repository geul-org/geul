//ff:func feature=resolve type=helper control=sequence
//ff:what listPath — `<Issuer>/<문서 ID>/versions.geul`, the document's version list address
package resolve

// listPath is the address of d's version list document, relative to
// `.well-known/geul/`.
func listPath(d DocKey) string {
	return d.Issuer + "/" + d.Doc + "/" + listFile
}
