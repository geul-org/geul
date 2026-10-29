//ff:func feature=pack type=decoder control=sequence
//ff:what plainDecoder — the shared zstd decoder (klauspost/compress) for frames without a dictionary, made on first use
package pack

import (
	"sync"

	"github.com/klauspost/compress/zstd"
)

var (
	plainOnce sync.Once
	plain     *zstd.Decoder
)

// plainDecoder returns the package decoder; DecodeAll is safe for concurrent
// use. NewReader fails only on bad options, and these are fixed.
func plainDecoder() *zstd.Decoder {
	plainOnce.Do(func() { plain, _ = zstd.NewReader(nil, zstd.WithDecoderConcurrency(0)) })
	return plain
}
