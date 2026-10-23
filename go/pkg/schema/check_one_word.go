//ff:func feature=schema type=rule control=sequence
//ff:what checkOneWord — D4: BODY and END are 1-word records
package schema

// checkOneWord applies D4 to a BODY or END record: no payload.
func checkOneWord(r rec, name string) []Fact {
	if r.Total != 1 {
		return []Fact{factAt("D4", Error, r, "Length", name+" 1 WORD", fmtWords(r.Total))}
	}
	return nil
}
