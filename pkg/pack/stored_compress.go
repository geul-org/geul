//ff:func feature=pack type=encoder control=sequence
//ff:what Stored.Compress — src as a raw-block zstd frame with content size and checksum
package pack

// Compress frames src without compressing it.
func (Stored) Compress(src []byte) ([]byte, error) {
	return zstdFrame(src), nil
}
