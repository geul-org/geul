//ff:func feature=pack type=decoder control=sequence
//ff:what Reader.unzstd — one zstd frame and nothing after: content size required, checksum when asked, dictionary from DICTS
package pack

import (
	"context"
	"fmt"
)

// unzstd decodes a frame. Frame_Content_Size is required (the decoder checks
// the result against it, and the checksum when present); a non-zero
// Dictionary_ID loads that DICTS entry. Any block type is read
// (klauspost/compress), so compressed records always open.
func (r *Reader) unzstd(ctx context.Context, frame []byte, needSum bool) ([]byte, error) {
	h, err := parseFrameHeader(frame)
	if err == nil && (!h.hasFCS || needSum && !h.checksum) {
		err = corrupt("zstd frame without content size or checksum")
	}
	var end int
	if err == nil {
		end, err = frameEnd(frame, h)
	}
	if err == nil && end != len(frame) {
		err = corrupt("bytes after the zstd frame")
	}
	if err != nil {
		return nil, err
	}
	d, err := r.decoder(ctx, h.dictID)
	if err != nil {
		return nil, err
	}
	out, err := d.DecodeAll(frame, make([]byte, 0, min(h.fcs, 8*uint64(len(frame))+4096)))
	if err != nil {
		return nil, fmt.Errorf("%w: zstd: %v", ErrCorrupt, err)
	}
	return out, nil
}
