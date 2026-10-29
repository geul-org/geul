//ff:type feature=pack type=model
//ff:what StreamRecord — one record as a Stream stores it: encoding, stored bytes, raw length and the content hash it also serves under h/
package pack

// StreamRecord is a record ready to store (Stream.Encode): Enc is encRaw or
// encZstd, Bytes the stored bytes, Raw the length before compression, and
// Hash the GEUL content hash when the bytes are a document (HashOK), under
// which the record also serves `h/<Hash>.geul` (flags bit0, §4.5).
type StreamRecord struct {
	Enc    byte
	Bytes  []byte
	Raw    int
	Hash   [32]byte
	HashOK bool
}
