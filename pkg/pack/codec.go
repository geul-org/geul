//ff:type feature=pack type=model
//ff:what Codec — the zstd frame seam: compress one record or block into a frame, decompress one frame
package pack

// Codec turns bytes into one zstd frame (RFC 8878) with Frame_Content_Size
// and Content_Checksum, and back. The built-in Stored writes raw-block frames
// and reads raw and RLE blocks; a real zstd codec (with dictionaries) plugs in
// here without pkg/pack depending on an outside module.
type Codec interface {
	// Compress returns one frame holding src.
	Compress(src []byte) ([]byte, error)
	// Decompress returns the content of one frame; dict is the frame's
	// dictionary (nil when its Dictionary_ID is 0).
	Decompress(frame, dict []byte) ([]byte, error)
}
