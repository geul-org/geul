//ff:type feature=pack type=model
//ff:what frameHeader — the parsed zstd frame header: length, content size, checksum flag, dictionary, window
package pack

// frameHeader is what a reader needs from a zstd frame header (RFC 8878 3.1.1.1).
type frameHeader struct {
	size     int // header bytes, magic included
	fcs      uint64
	hasFCS   bool
	checksum bool
	dictID   uint32
	window   uint64
}
