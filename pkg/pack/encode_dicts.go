//ff:func feature=pack type=encoder control=iteration dimension=1
//ff:what encodeDicts — the DICTS section body
package pack

import "encoding/binary"

// encodeDicts writes DICTS.
func encodeDicts(ds []dictEntry) []byte {
	out := binary.LittleEndian.AppendUint32(nil, uint32(len(ds)))
	for _, d := range ds {
		out = binary.LittleEndian.AppendUint32(out, d.id)
		out = binary.LittleEndian.AppendUint32(out, d.loc.Obj)
		out = binary.LittleEndian.AppendUint64(out, d.loc.Off)
		out = binary.LittleEndian.AppendUint32(out, uint32(d.loc.Len))
		out = append(out, d.hash[:]...)
	}
	return out
}
