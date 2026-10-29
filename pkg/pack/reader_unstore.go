//ff:func feature=pack type=decoder control=selection
//ff:what Reader.unstore — stored bytes by encoding: raw as they are, zstd through the codec
package pack

import "context"

// unstore restores stored bytes; needSum demands a zstd Content_Checksum
// (blocks, inventories).
func (r *Reader) unstore(ctx context.Context, enc byte, b []byte, needSum bool) ([]byte, error) {
	switch enc {
	case encRaw:
		return b, nil
	case encZstd:
		return r.unzstd(ctx, b, needSum)
	default:
		return nil, corrupt("unknown encoding")
	}
}
