//ff:func feature=pack type=encoder control=sequence
//ff:what Zstd.Compress — src as one compressed zstd frame (content size, checksum, d's ID); short input as a raw-block frame
package pack

// Compress compresses src. Input shorter than 256 bytes stays a raw-block
// frame: klauspost writes such frames multi-segment without a content size,
// which §4.1 requires, and they hardly shrink anyway.
func (z *Zstd) Compress(src []byte, d *Dict) ([]byte, error) {
	if len(src) < minZstd {
		return zstdFrame(src), nil
	}
	e, err := z.encoder(d)
	if err != nil {
		return nil, err
	}
	return e.EncodeAll(src, nil), nil
}
