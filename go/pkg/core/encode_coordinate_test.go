//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeCoordinate — the docs' Berlin (COORDINATE-EARTH) and Mars (COORDINATE-GEO) words
package core

import (
	"slices"
	"testing"
)

func TestEncodeCoordinate(t *testing.T) {
	ws, _ := EncodeRecord(EncodeCoordinate(Coordinate{Lat: 52.516666666667, Lon: 13.383333333333, U: 2172487}))
	if !slices.Equal(ws, []Word{0x0D04, 0x404A, 0x4222, 0x2222, 0x2251, 0x402A, 0xC444, 0x4444, 0x4389, 0x0021, 0x2647}) {
		t.Fatalf("earth: %04X", ws)
	}
	mars := Record{Kind: KindReference, Codec: RefWikidataQID << 3, Payload: []Word{111}}
	ws, _ = EncodeRecord(EncodeCoordinate(Coordinate{Globe: &mars, Lat: 4.5, Lon: 137.4, U: UncertaintyUnknown}))
	if !slices.Equal(ws, []Word{0x0E05, 0x00A0, 0x006F, 0x4012, 0, 0, 0, 0x4061, 0x2CCC, 0xCCCC, 0xCCCD, 0xFFFF, 0xFFFF}) {
		t.Fatalf("geo: %04X", ws)
	}
}
