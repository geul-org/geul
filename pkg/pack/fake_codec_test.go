//ff:type feature=pack type=model
//ff:what fakeCodec — a Codec that fails or returns a fixed frame, for writer branches
package pack

// fakeCodec returns out (or err) from Compress and decodes with Stored.
type fakeCodec struct {
	out []byte
	err error
}
