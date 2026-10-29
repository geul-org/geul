//ff:func feature=pack type=encoder control=sequence
//ff:what encodeBlock — a directory block in its byte form: kind, varint n, key columns, location columns
package pack

import "encoding/binary"

// encodeBlock serializes b (§4.7) before any zstd framing.
func encodeBlock(b block) []byte {
	out := []byte{b.kind}
	out = binary.AppendUvarint(out, uint64(len(b.keys)))
	if isPKind(b.kind) {
		out = appendPKeys(out, b.keys)
	} else {
		out = appendHKeys(out, b.keys)
	}
	return appendLocs(out, b.locs)
}
