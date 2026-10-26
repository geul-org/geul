//ff:func feature=core type=encoder control=iteration dimension=1
//ff:what EncodeQuantity — a QUANTITY Value from its amount, unit and bound records
package core

// EncodeQuantity builds the QUANTITY Value of q: amount, unit, lower and
// upper records in that order, each whole (a left-out one as the one-word
// Absent record). It fails when a record does not encode.
func EncodeQuantity(q Quantity) (Record, error) {
	var p []Word
	for _, r := range []Record{q.Amount, q.Unit, q.Lower, q.Upper} {
		ws, err := EncodeRecord(r)
		if err != nil {
			return Record{}, err
		}
		p = append(p, ws...)
	}
	return Record{Kind: KindValue, Codec: CodecQuantity, Payload: p}, nil
}
