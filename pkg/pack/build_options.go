//ff:type feature=pack type=model
//ff:what BuildOptions — pack writer parameters: object size, block sizes, H layer limit, codec, clock, verification
package pack

import (
	"context"
	"time"
)

// BuildOptions tune a new generation (§6). Zero values take the defaults.
// Verify re-opens the new manifest before latest.txt moves and checks every
// put path's bytes, every deleted path's absence and, for a Full batch, the
// path set. Check runs after that on the new generation. Fresh rewrites every
// record into new objects (compaction): the new OBJECTS lists only them.
// Codec nil follows the set: Zstd when its current generation has a
// dictionary or a compressed record, Stored otherwise (and for a new set).
// Batch.Get may be called from several goroutines when Codec is Zstd.
type BuildOptions struct {
	PackSize int64
	PLeaf    int
	HLeaf    int
	HFanout  int
	HRoot    int
	HLayers  int
	Codec    Codec
	Now      func() time.Time
	Verify   bool
	Check    func(ctx context.Context, r *Reader) error
	Fresh    bool
}
