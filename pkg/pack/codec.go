//ff:type feature=pack type=model
//ff:what Codec — how a writer frames bytes: one zstd frame per record, block or inventory, with an optional dictionary
package pack

// Codec turns bytes into one zstd frame (RFC 8878) with Frame_Content_Size and
// Content_Checksum. Stored writes raw blocks and leaves records raw; Zstd
// compresses (klauspost/compress). BuildOptions.Codec nil follows the set:
// Zstd when its current generation has compressed records or a dictionary,
// Stored otherwise. Reading needs no codec.
type Codec interface {
	// Compress returns one frame holding src; d is nil or the dictionary new
	// records use, whose ID the frame then names.
	Compress(src []byte, d *Dict) ([]byte, error)
}
