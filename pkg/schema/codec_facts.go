//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what codecFacts — layer 2: every Codec that matches one record checks it
package schema

// codecFacts runs each Codec that matches r.
func codecFacts(r rec, codecs []Codec) []Fact {
	var fs []Fact
	for _, c := range codecs {
		if c.Match(r.R) {
			fs = append(fs, c.Check(r.R, Where{Word: r.Word, Record: r.Index})...)
		}
	}
	return fs
}
