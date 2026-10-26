//ff:func feature=schema type=rule control=sequence
//ff:what TestQuantityDecimals — minimal amount and bounds pass, Absent bounds are skipped, a non-minimal upper bound fails
package schema

import (
	"math/big"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestQuantityDecimals(t *testing.T) {
	absent := core.Record{Kind: core.KindReference}
	d := core.EncodeSignedDecimalBig(0, big.NewInt(5))
	if err := quantityDecimals(core.Quantity{Amount: d, Lower: d, Upper: absent}); err != nil {
		t.Fatal(err)
	}
	loose := core.Record{Kind: core.KindValue, Codec: core.CodecSignedDecimal, Payload: []core.Word{0, 0, 0, 1}}
	if quantityDecimals(core.Quantity{Amount: d, Lower: absent, Upper: loose}) == nil {
		t.Fatal("non-minimal upper: want error")
	}
}
