//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what parseIssuers — an Issuer table: u32 count, entries of width (1·2·4·16 words) and big-endian ID bytes
package pack

import "encoding/hex"

// parseIssuers decodes an Issuer table into canonical hex; duplicates and
// other widths are ErrCorrupt. Trailing bytes are left in c.
func parseIssuers(c *cursor) ([]string, error) {
	n := c.u32()
	if c.bad || uint64(n)*3 > uint64(len(c.b)) {
		return nil, corrupt("Issuer table")
	}
	out := make([]string, n)
	seen := map[string]bool{}
	for i := range out {
		w := c.u8()
		id := "0x" + hex.EncodeToString(c.take(uint64(w)*2))
		if c.bad || w != 1 && w != 2 && w != 4 && w != 16 || seen[id] {
			return nil, corrupt("Issuer table entry")
		}
		out[i], seen[id] = id, true
	}
	return out, nil
}
