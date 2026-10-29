//ff:func feature=pack type=encoder control=iteration dimension=1
//ff:what appendIssuers — an Issuer table: u32 count, then width in words and big-endian ID bytes
package pack

import (
	"encoding/binary"
	"encoding/hex"
)

// appendIssuers writes ids (canonical hex) as an Issuer table.
func appendIssuers(out []byte, ids []string) []byte {
	out = binary.LittleEndian.AppendUint32(out, uint32(len(ids)))
	for _, id := range ids {
		b, _ := hex.DecodeString(id[2:])
		out = append(out, byte(len(b)/2))
		out = append(out, b...)
	}
	return out
}
