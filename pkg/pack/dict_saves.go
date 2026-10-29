//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what dictSaves — whether d stores the test records smaller in total than no dictionary (a record never stores bigger than raw)
package pack

// dictSaves compresses each test record with and without d and compares the
// stored sizes (min of compressed and raw).
func dictSaves(z *Zstd, d *Dict, test [][]byte) bool {
	var plain, with int
	for _, s := range test {
		a, _ := z.Compress(s, nil) // fails only with a dictionary
		b, err := z.Compress(s, d)
		if err != nil {
			return false
		}
		plain += min(len(a), len(s))
		with += min(len(b), len(s))
	}
	return with < plain
}
