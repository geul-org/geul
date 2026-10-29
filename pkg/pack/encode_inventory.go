//ff:func feature=pack type=encoder control=iteration dimension=1
//ff:what encodeInventory — the inventory before zstd framing: Issuer table, path table, entries, deleted hashes
package pack

import "encoding/binary"

// encodeInventory serializes inv (§4.5).
func encodeInventory(inv inventory) []byte {
	out := appendIssuers(nil, inv.issuers)
	out = binary.LittleEndian.AppendUint32(out, uint32(len(inv.paths)))
	for _, p := range inv.paths {
		out = binary.LittleEndian.AppendUint16(out, uint16(len(p)))
		out = append(out, p...)
	}
	out = binary.LittleEndian.AppendUint64(out, uint64(len(inv.entries)))
	for _, e := range inv.entries {
		out = append(out, e.kind, e.enc)
		out = binary.LittleEndian.AppendUint16(out, e.flags)
		out = binary.LittleEndian.AppendUint32(out, e.ref)
		out = binary.LittleEndian.AppendUint64(out, e.doc)
		out = binary.LittleEndian.AppendUint32(out, e.slot)
		out = binary.LittleEndian.AppendUint32(out, e.n)
		out = binary.LittleEndian.AppendUint64(out, e.off)
	}
	out = binary.LittleEndian.AppendUint64(out, uint64(len(inv.deleted)))
	for _, h := range inv.deleted {
		out = append(out, h...)
	}
	return out
}
