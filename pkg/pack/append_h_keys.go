//ff:func feature=pack type=encoder control=iteration dimension=1
//ff:what appendHKeys — C1 of H keys, the 32-byte hashes as they are
package pack

// appendHKeys writes the H key column.
func appendHKeys(out []byte, keys []string) []byte {
	for _, k := range keys {
		out = append(out, k...)
	}
	return out
}
