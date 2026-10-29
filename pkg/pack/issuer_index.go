//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what issuerIndex — Issuer hex → its number (position in the table)
package pack

// issuerIndex maps each Issuer to its position.
func issuerIndex(ids []string) map[string]uint32 {
	m := make(map[string]uint32, len(ids))
	for i, id := range ids {
		m[id] = uint32(i)
	}
	return m
}
