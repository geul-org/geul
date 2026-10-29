//ff:func feature=pack type=decoder control=sequence
//ff:what newDictDecoder — a zstd decoder for one DICTS dictionary: formatted (its own ID must be the DICTS ID) or raw content
package pack

import (
	"fmt"

	"github.com/klauspost/compress/zstd"
)

// newDictDecoder loads dictionary b under id.
func newDictDecoder(id uint32, b []byte) (*zstd.Decoder, error) {
	opt := zstd.WithDecoderDictRaw(id, b)
	if fid, ok := dictID(b); ok {
		if fid != id {
			return nil, corrupt("zstd dictionary ID differs from its DICTS entry")
		}
		opt = zstd.WithDecoderDicts(b)
	}
	d, err := zstd.NewReader(nil, zstd.WithDecoderConcurrency(0), opt)
	if err != nil {
		return nil, fmt.Errorf("%w: zstd dictionary: %v", ErrCorrupt, err)
	}
	return d, nil
}
