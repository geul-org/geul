//ff:func feature=core type=helper control=iteration dimension=1
//ff:what stripValidation — the words of a well-formed record run with validation-class Meta records removed
package core

// stripValidation copies the records of ws, a run already checked by
// SplitDocument, leaving out every validation-class Meta record (class 100,
// codecs 0x40-0x4F).
func stripValidation(ws []Word) []Word {
	out := make([]Word, 0, len(ws))
	for i := 0; i < len(ws); {
		r, m, _ := DecodeRecord(ws[i:])
		if r.Kind != KindMeta || r.Codec>>4 != 4 {
			out = append(out, ws[i:i+m]...)
		}
		i += m
	}
	return out
}
