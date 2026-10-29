//ff:func feature=pack type=encoder control=sequence
//ff:what fakeCodec.Compress — the fixed result
package pack

// Compress returns c.out, c.err.
func (c fakeCodec) Compress([]byte) ([]byte, error) {
	return c.out, c.err
}
