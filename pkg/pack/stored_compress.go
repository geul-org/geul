//ff:func feature=pack type=encoder control=sequence
//ff:what Stored.Compress — src as a raw-block zstd frame with content size and checksum (no dictionary)
package pack

// Compress frames src without compressing it; d is ignored.
func (Stored) Compress(src []byte, _ *Dict) ([]byte, error) {
	return zstdFrame(src), nil
}
