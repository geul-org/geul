//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeCoordinate — Earth and Geo round trip; another codec, a missing globe and a wrong width fail
package core

import "testing"

func TestDecodeCoordinate(t *testing.T) {
	in := Coordinate{Lat: -33.5, Lon: 151.25, U: 44000}
	if c, err := DecodeCoordinate(EncodeCoordinate(in)); err != nil || c.Lat != in.Lat || c.Lon != in.Lon || c.U != in.U || c.Globe != nil {
		t.Fatalf("earth: %+v %v", c, err)
	}
	moon := Record{Kind: KindReference, Codec: RefWikidataQID << 3, Payload: []Word{405}}
	if c, err := DecodeCoordinate(EncodeCoordinate(Coordinate{Globe: &moon, Lat: 1, Lon: 359.5})); err != nil || c.Globe.Payload[0] != 405 || c.Lon != 359.5 {
		t.Fatalf("geo: %+v %v", c, err)
	}
	for i, r := range []Record{
		EncodeUintValue(1),
		{Kind: KindValue, Codec: CodecCoordinateGeo, Payload: make([]Word, 10)},
		{Kind: KindValue, Codec: CodecCoordinateGeo, Payload: []Word{0x08C0}},
		{Kind: KindValue, Codec: CodecCoordinateEarth, Payload: make([]Word, 9)},
	} {
		if _, err := DecodeCoordinate(r); err == nil {
			t.Fatalf("case %d: want error", i)
		}
	}
}
