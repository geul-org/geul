//ff:type feature=pack type=resolver
//ff:what Archive — an archive kept as pack sets, read as an fs.FS rooted at `.well-known/geul/` and written by staging changes until Commit
package pack

import "sync"

// Archive serves the paths of one or more pack sets as a read-only fs.FS
// (ReadFile, ReadDir, Stat, Open) and implements resolve.Store: WriteFile and
// Remove stage changes that reads see at once and Commit turns into one new
// generation per changed set. Sets live in the `.well-known/geul/` directory:
// `latest.txt` at its top (one set for everything) and/or `<Issuer>/latest.txt`
// (one set per Issuer). A remote Archive (OpenRemote) cannot be written.
type Archive struct {
	dir  string // the `.well-known/geul/` directory; "" when remote
	o    BuildOptions
	mu   sync.Mutex
	sets []*packSet // read order: the top set first, then Issuer sets in Dir order
	puts map[string][]byte
	dels map[string]bool
	idx  []string // sorted live paths, built on first listing
}
