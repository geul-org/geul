//ff:func feature=pack type=encoder control=iteration dimension=1
//ff:what encodeObjects — the OBJECTS section body
package pack

import "encoding/binary"

// encodeObjects writes OBJECTS.
func encodeObjects(objs []object) []byte {
	out := binary.LittleEndian.AppendUint32(nil, uint32(len(objs)))
	for _, o := range objs {
		out = append(out, o.kind, 0, 0, 0, 0, 0, 0, 0)
		out = binary.LittleEndian.AppendUint64(out, o.size)
		out = append(out, o.hash[:]...)
	}
	return out
}
