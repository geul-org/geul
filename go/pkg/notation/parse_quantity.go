//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what parseQuantity — QUANTITY from four record tokens: amount, unit, lower, upper (_ for a part left out)
package notation

import (
	"errors"

	"github.com/geul-org/geul/go/pkg/core"
)

// parseQuantity reads the four parts (parseRecord), encodes them and checks
// the result as a QUANTITY (core.DecodeQuantity: amount and bounds decimals,
// unit a Reference).
func parseQuantity(args []string, ctx Context) (core.Record, error) {
	if len(args) != 4 {
		return core.Record{}, errors.New("geul/notation: QUANTITY takes amount, unit, lower, upper")
	}
	var rs [4]core.Record
	for i, a := range args {
		r, err := parseRecord(a, ctx)
		if err != nil {
			return core.Record{}, err
		}
		rs[i] = r
	}
	r, _ := core.EncodeQuantity(core.Quantity{Amount: rs[0], Unit: rs[1], Lower: rs[2], Upper: rs[3]})
	_, err := core.DecodeQuantity(r)
	return r, err
}
