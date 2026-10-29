//ff:func feature=pack type=encoder control=iteration dimension=1
//ff:what zstdFrame — src as one zstd frame of raw blocks: 128 KiB window, 8-byte content size, XXH64 checksum
package pack

import "encoding/binary"

// zstdFrame wraps src in a valid zstd frame any decoder reads: descriptor
// 0xC4 (8-byte Frame_Content_Size, Content_Checksum), window 128 KiB, raw
// blocks of at most 128 KiB, then the low 32 bits of XXH64.
func zstdFrame(src []byte) []byte {
	out := binary.LittleEndian.AppendUint32(nil, zstdMagic)
	out = append(out, 0xC4, 0x38)
	out = binary.LittleEndian.AppendUint64(out, uint64(len(src)))
	rest := src
	for {
		n := min(len(rest), maxBlock)
		last := n == len(rest)
		out = appendRawBlock(out, rest[:n], last)
		rest = rest[n:]
		if last {
			break
		}
	}
	return binary.LittleEndian.AppendUint32(out, uint32(xxh64(src)))
}
