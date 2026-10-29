//ff:func feature=pack type=decoder control=sequence
//ff:what Stored.Decompress — the content of one zstd frame of raw/RLE blocks, checksum checked, nothing after
package pack

// Decompress decodes one frame; the dictionary is not needed for raw and RLE
// blocks. Trailing bytes after the frame are ErrCorrupt.
func (Stored) Decompress(frame, _ []byte) ([]byte, error) {
	h, err := parseFrameHeader(frame)
	if err != nil {
		return nil, err
	}
	out, rest, err := zstdBlocks(frame[h.size:], h.window)
	if err == nil && h.checksum {
		rest, err = checkXXH(out, rest)
	}
	if err == nil && len(rest) != 0 {
		err = corrupt("bytes after the zstd frame")
	}
	if err != nil {
		return nil, err
	}
	return out, nil
}
