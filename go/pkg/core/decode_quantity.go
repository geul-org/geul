//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what DecodeQuantity — the amount, unit and bounds of a QUANTITY Value, checking each record's form and that nothing is left over
package core

import (
	"errors"
	"fmt"
)

// DecodeQuantity reads the four records of a QUANTITY Value r
// (value/numeric/QUANTITY): the amount must be a SIGNED-DECIMAL Value, the
// unit a Reference (or Absent), each bound a SIGNED-DECIMAL Value or Absent,
// and the payload must end right after the fourth record.
func DecodeQuantity(r Record) (Quantity, error) {
	if r.Kind != KindValue || r.Codec != CodecQuantity {
		return Quantity{}, fmt.Errorf("geul/core: expected a QUANTITY Value, got kind %d codec 0x%02X", r.Kind, r.Codec)
	}
	var rs [4]Record
	ws := r.Payload
	for i := range rs {
		rec, n, err := DecodeRecord(ws)
		if err != nil {
			return Quantity{}, fmt.Errorf("geul/core: QUANTITY record %d: %w", i+1, err)
		}
		rs[i], ws = rec, ws[n:]
	}
	if len(ws) != 0 {
		return Quantity{}, errors.New("geul/core: QUANTITY has words after the upper bound")
	}
	if err := quantityForm(rs); err != nil {
		return Quantity{}, err
	}
	return Quantity{Amount: rs[0], Unit: rs[1], Lower: rs[2], Upper: rs[3]}, nil
}
