//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what runCodecs — layer 2 over every framed record
package schema

// runCodecs applies the layer-2 Codecs to every framed record.
func runCodecs(recs []rec, codecs []Codec) []Fact {
	var fs []Fact
	for _, r := range recs {
		fs = append(fs, codecFacts(r, codecs)...)
	}
	return fs
}
