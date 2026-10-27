//ff:type feature=resolve type=model
//ff:what Config — where to fetch from: the source list and the private patterns
package resolve

// Config is the fetching side's location setting (GEUL-ORG 위치 해석). Proxy
// is the ordered source list (GEULPROXY) and Private the documents never sent
// to public services (GEULPRIVATE). GEULLOG, GEULNOLOG, GEULINSECURE and the
// workspace are not implemented yet.
type Config struct {
	Proxy   []Entry
	Private []Pattern
}
