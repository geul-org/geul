//ff:func feature=pack type=decoder control=sequence
//ff:what fakeCodec.Decompress — Stored decoding
package pack

// Decompress decodes like Stored.
func (c fakeCodec) Decompress(frame, dict []byte) ([]byte, error) {
	return Stored{}.Decompress(frame, dict)
}
