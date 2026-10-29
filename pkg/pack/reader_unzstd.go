//ff:func feature=pack type=decoder control=sequence
//ff:what Reader.unzstd — one zstd frame through the codec: content size required, checksum when asked, dictionary from DICTS
package pack

import "context"

// unzstd decodes a frame. Frame_Content_Size is required and must equal the
// result length; a non-zero Dictionary_ID loads that DICTS entry.
func (r *Reader) unzstd(ctx context.Context, frame []byte, needSum bool) ([]byte, error) {
	h, err := parseFrameHeader(frame)
	if err == nil && (!h.hasFCS || needSum && !h.checksum) {
		err = corrupt("zstd frame without content size or checksum")
	}
	var dict []byte
	if err == nil && h.dictID != 0 {
		dict, err = r.dict(ctx, h.dictID)
	}
	if err != nil {
		return nil, err
	}
	out, err := r.codec.Decompress(frame, dict)
	if err == nil && uint64(len(out)) != h.fcs {
		err = corrupt("zstd content size")
	}
	if err != nil {
		return nil, err
	}
	return out, nil
}
