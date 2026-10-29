//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what writer.dictPays — whether a dictionary saves more than its own size: held-out samples with and without it, scaled to all records
package pack

// dictPays compares stored sizes (a record never stores bigger than raw) on
// test and scales the saving by the bytes of every record to write.
func (w *writer) dictPays(z *Zstd, d *Dict, test [][]byte) bool {
	var raw, plain, with float64
	for _, s := range test {
		a, _ := z.Compress(s, nil) // fails only with a dictionary
		b, err := z.Compress(s, d)
		if err != nil {
			return false
		}
		raw += float64(len(s))
		plain += float64(min(len(a), len(s)))
		with += float64(min(len(b), len(s)))
	}
	return raw > 0 && (plain-with)*float64(w.recordBytes())/raw > float64(len(d.Bytes))
}
