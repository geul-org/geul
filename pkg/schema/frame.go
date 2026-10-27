//ff:func feature=schema type=decoder control=iteration dimension=1
//ff:what frame — split the words into records, stopping where a length cannot be trusted
package schema

import "github.com/geul-org/gogeul/pkg/core"

// frame walks ws record by record using each Length, without a codebook.
func frame(ws []core.Word) ([]rec, []Fact) {
	var recs []rec
	var fs []Fact
	for at, i := 0, 0; at < len(ws); i++ {
		r, f, ok := readRec(ws, at, i)
		fs = append(fs, f...)
		if !ok {
			return recs, fs
		}
		recs = append(recs, r)
		at += r.Total
	}
	return recs, fs
}
