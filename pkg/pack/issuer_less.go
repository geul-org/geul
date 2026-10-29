//ff:func feature=pack type=helper control=sequence
//ff:what issuerLess — Issuer order of resolve.Dir: shorter IDs first, then by hex
package pack

// issuerLess orders canonical Issuer hex by width, then value.
func issuerLess(a, b string) bool {
	return len(a) < len(b) || len(a) == len(b) && a < b
}
