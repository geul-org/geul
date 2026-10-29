//ff:func feature=pack type=encoder control=iteration dimension=1
//ff:what encodeNamed — the NAMED section body
package pack

import "encoding/binary"

// encodeNamed writes NAMED; ns must be sorted by path.
func encodeNamed(ns []namedEntry) []byte {
	out := binary.LittleEndian.AppendUint32(nil, uint32(len(ns)))
	for _, n := range ns {
		out = binary.LittleEndian.AppendUint16(out, uint16(len(n.path)))
		out = append(out, n.path...)
		out = append(out, n.loc.Enc, 0)
		out = binary.LittleEndian.AppendUint32(out, n.loc.Obj)
		out = binary.LittleEndian.AppendUint64(out, n.loc.Off)
		out = binary.LittleEndian.AppendUint32(out, uint32(n.loc.Len))
	}
	return out
}
