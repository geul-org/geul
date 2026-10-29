//ff:func feature=pack type=encoder control=iteration dimension=1
//ff:what encodeLayers — a P_LAYERS or H_LAYERS section body (newest layer first)
package pack

import "encoding/binary"

// encodeLayers writes the layers with their stored roots.
func encodeLayers(ls []layer) []byte {
	out := binary.LittleEndian.AppendUint32(nil, uint32(len(ls)))
	for _, l := range ls {
		out = binary.LittleEndian.AppendUint64(out, l.count)
		out = append(out, l.enc, 0, 0, 0)
		out = binary.LittleEndian.AppendUint32(out, uint32(len(l.root)))
		out = append(out, l.root...)
	}
	return out
}
