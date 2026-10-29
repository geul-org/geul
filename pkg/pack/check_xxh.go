//ff:func feature=pack type=rule control=sequence
//ff:what checkXXH — verify the 4-byte Content_Checksum after a frame's blocks; returns the bytes after it
package pack

import "encoding/binary"

// checkXXH compares the low 32 bits of XXH64(content) with the stored checksum.
func checkXXH(content, rest []byte) ([]byte, error) {
	if len(rest) < 4 || binary.LittleEndian.Uint32(rest) != uint32(xxh64(content)) {
		return nil, corrupt("zstd content checksum")
	}
	return rest[4:], nil
}
