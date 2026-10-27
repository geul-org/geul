//ff:func feature=core type=encoder control=iteration dimension=1
//ff:what EncodeFraction — a FRACTION Value from its numerator and denominator records
package core

// EncodeFraction builds the FRACTION Value of f: the numerator record, then
// the denominator record, each whole. It fails when a record does not encode;
// the codecs are DecodeFraction's check.
func EncodeFraction(f Fraction) (Record, error) {
	var p []Word
	for _, r := range []Record{f.Numerator, f.Denominator} {
		ws, err := EncodeRecord(r)
		if err != nil {
			return Record{}, err
		}
		p = append(p, ws...)
	}
	return Record{Kind: KindValue, Codec: CodecFraction, Payload: p}, nil
}
