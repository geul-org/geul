//ff:type feature=pack type=model
//ff:what inventory — a data pack's self-description: Issuer table, path table, entries, deleted h/ hashes
package pack

// inventory lists what a data pack adds and deletes, in application order (§4.5, §7.1).
type inventory struct {
	issuers []string
	paths   []string
	entries []invEntry
	deleted []string // 32-byte content hashes whose h/ paths this pack deletes
}
