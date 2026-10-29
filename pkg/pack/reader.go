//ff:type feature=pack type=resolver
//ff:what Reader — one generation opened from a Store: a resolve.Source plus key lookups, listings and verification
package pack

import (
	"sync"

	"github.com/klauspost/compress/zstd"
)

// Reader answers path → bytes for one manifest (§5.1). It is safe for
// concurrent use; directory blocks, dictionaries and their decoders are cached. Reading needs
// no codec: raw, RLE and compressed zstd blocks all decode.
type Reader struct {
	s    Store
	m    *manifest
	name string
	inum map[string]uint32

	mu     sync.Mutex
	blocks map[Loc]block
	dicts  map[uint32][]byte
	decs   map[uint32]*zstd.Decoder
	limit  int
}
