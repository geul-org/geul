//ff:func feature=pack type=rule control=iteration dimension=1
//ff:what Reader.objectSum — the SHA-256 of an object's size bytes, read in hashChunk ranges
package pack

import (
	"context"
	"crypto/sha256"
)

// objectSum hashes the first o.size bytes of o.
func (r *Reader) objectSum(ctx context.Context, o object) ([32]byte, error) {
	h := sha256.New()
	for off := uint64(0); off < o.size; off += hashChunk {
		b, err := r.readRange(ctx, o, off, min(hashChunk, o.size-off))
		if err != nil {
			return [32]byte{}, err
		}
		h.Write(b)
	}
	var sum [32]byte
	copy(sum[:], h.Sum(nil))
	return sum, nil
}
