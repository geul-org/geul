//ff:type feature=pack type=resolver
//ff:what Reader — one generation opened from a Store: a resolve.Source plus key lookups, listings and verification
package pack

import "sync"

// Reader answers path → bytes for one manifest (§5.1). It is safe for
// concurrent use; directory blocks and dictionaries are cached.
type Reader struct {
	s     Store
	codec Codec
	m     *manifest
	name  string
	inum  map[string]uint32

	mu     sync.Mutex
	blocks map[Loc]block
	dicts  map[uint32][]byte
	limit  int
}
