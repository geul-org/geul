//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckQuantity — a good QUANTITY passes; a wrong layout and a non-minimal coefficient are V9
package schema

import (
	"math/big"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCheckQuantity(t *testing.T) {
	absent := core.Record{Kind: core.KindReference}
	good, _ := core.EncodeQuantity(core.Quantity{Amount: core.EncodeSignedDecimalBig(-2, big.NewInt(120)), Unit: absent, Lower: absent, Upper: absent})
	if fs := checkQuantity(recOf(good)); fs != nil {
		t.Fatalf("good: %+v", fs)
	}
	if rulesOf(checkQuantity(recOf(core.Record{Kind: core.KindValue, Codec: core.CodecQuantity, Payload: []core.Word{0}}))) != "V9" {
		t.Fatal("layout: want V9")
	}
	loose := core.Record{Kind: core.KindValue, Codec: core.CodecSignedDecimal, Payload: []core.Word{0, 0, 0, 1}}
	bad, _ := core.EncodeQuantity(core.Quantity{Amount: loose, Unit: absent, Lower: absent, Upper: absent})
	if rulesOf(checkQuantity(recOf(bad))) != "V9" {
		t.Fatal("non-minimal: want V9")
	}
}
