//ff:func feature=pack type=encoder control=sequence
//ff:what newEncoder — a klauspost zstd encoder: level, concurrency, checksum, and the dictionary (formatted or raw content) when given
package pack

import (
	"fmt"

	"github.com/klauspost/compress/zstd"
)

// newEncoder makes an encoder whose frames carry Content_Checksum and, with
// d, d's ID. A formatted dictionary must carry d.ID itself.
func newEncoder(level zstd.EncoderLevel, n int, d *Dict) (*zstd.Encoder, error) {
	opts := []zstd.EOption{zstd.WithEncoderLevel(level), zstd.WithEncoderConcurrency(n), zstd.WithEncoderCRC(true)}
	if d != nil {
		opts = append(opts, zstd.WithEncoderDictRaw(d.ID, d.Bytes))
	}
	if fid, ok := dictID(dictBytes(d)); ok {
		if fid != d.ID {
			return nil, fmt.Errorf("geul/pack: dictionary %d carries ID %d", d.ID, fid)
		}
		opts[len(opts)-1] = zstd.WithEncoderDict(d.Bytes)
	}
	return zstd.NewWriter(nil, opts...)
}
