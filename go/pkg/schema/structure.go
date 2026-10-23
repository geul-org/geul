//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what structure — D1~D4 over the framed records, splitting HEAD and BODY
package schema

// structure checks the document structure (D1~D4) and returns the HEAD and BODY
// records. words is the input length in words.
func structure(recs []rec, words int) (docParts, []Fact) {
	var parts docParts
	if len(recs) == 0 {
		return parts, []Fact{{Rule: "D1", Severity: Error, Expected: "DOC-* 레코드", Actual: "레코드 없음"}}
	}
	fs := checkDocOpen(recs[0], words)
	st := &placeState{}
	for _, r := range recs[1:] {
		fs = append(fs, place(st, &parts, r)...)
	}
	if !st.ended {
		last := recs[len(recs)-1]
		fs = append(fs, factAt("D2", Error, last, "", "END로 끝나는 문서", "END 없음"))
	}
	return parts, fs
}
