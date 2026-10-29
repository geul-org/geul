//ff:func feature=pack type=encoder control=iteration dimension=1
//ff:what encodeManifest — a `.gpm`: 64-byte header, section table, then the six section bodies
package pack

import (
	"encoding/binary"
)

// encodeManifest serializes m (§4.3) with sections in kind order.
func encodeManifest(m *manifest) []byte {
	bodies := [][]byte{
		encodeObjects(m.objects), appendIssuers(nil, m.issuers), encodeDicts(m.dicts),
		encodeNamed(m.named), encodeLayers(m.pLayers), encodeLayers(m.hLayers),
	}
	out := []byte(manifestMagic)
	out = binary.LittleEndian.AppendUint16(out, formatVersion)
	out = binary.LittleEndian.AppendUint16(out, 0)
	out = binary.LittleEndian.AppendUint32(out, uint32(len(bodies)))
	out = binary.LittleEndian.AppendUint64(out, m.gen)
	out = binary.LittleEndian.AppendUint64(out, uint64(m.created))
	out = append(out, m.prev[:]...)
	off := uint64(headerSize + 24*len(bodies))
	for i, b := range bodies {
		out = binary.LittleEndian.AppendUint16(out, uint16(i+1))
		out = append(out, 0, 0, 0, 0, 0, 0)
		out = binary.LittleEndian.AppendUint64(out, off)
		out = binary.LittleEndian.AppendUint64(out, uint64(len(b)))
		off += uint64(len(b))
	}
	for _, b := range bodies {
		out = append(out, b...)
	}
	return out
}
