//ff:func feature=resolve type=helper control=sequence
//ff:what versionPath — `<Issuer>/<문서 ID>/<버전 ID>.geul`, the canonical address of one version
package resolve

// versionPath is the immutable canonical address of version v of d, relative
// to `.well-known/geul/`.
func versionPath(d DocKey, v uint32) string {
	return d.Issuer + "/" + d.Doc + "/" + FormatVersion(v) + ".geul"
}
