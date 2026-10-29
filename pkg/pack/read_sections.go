//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what readSections — the manifest section table: one body per kind, unknown kinds below 0x8000 rejected
package pack

import "fmt"

// readSections reads n 24-byte section entries from c and slices each body
// out of the whole manifest b. Kinds ≥ 0x8000 are skipped.
func readSections(c *cursor, n uint32, b []byte) (map[uint16][]byte, error) {
	secs := map[uint16][]byte{}
	for i := uint32(0); i < n; i++ {
		kind, res := c.u16(), c.take(6)
		off, ln := c.u64(), c.u64()
		_, dup := secs[kind]
		if c.bad || dup || kind == 0 || kind > secHLayers && kind < 0x8000 || !zeros(res) || off > uint64(len(b)) || ln > uint64(len(b))-off {
			return nil, corrupt(fmt.Sprintf("manifest section %d", i))
		}
		secs[kind] = b[off : off+ln]
	}
	return secs, nil
}
