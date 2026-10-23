package resolve

const (
	// DefaultProxy is the GEULPROXY used when the variable is empty (GEUL-ORG 출처 목록).
	DefaultProxy = "https://geul.org,registry"
	// geulOrg is the base of the public geul.org source, skipped for private documents.
	geulOrg = "https://geul.org"
	// wellKnown is the path every source serves documents under (RFC 8615).
	wellKnown = ".well-known/geul"
	// listFile is a document's version list document (edge/PUBLICATION).
	listFile = "versions.geul"
)
