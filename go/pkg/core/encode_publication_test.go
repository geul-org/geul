//ff:func feature=core type=encoder control=iteration dimension=1
//ff:what TestEncodePublication — each codec round-trips through DecodePublication
package core

import (
	"reflect"
	"testing"
)

func TestEncodePublication(t *testing.T) {
	doc := []Word{0, 0, 0, 5}
	for _, p := range []Publication{
		{Codec: EdgePublication, EdgeID: 1, Doc: doc, Version: 2, Hash: [32]byte{31: 0xAB}},
		{Codec: EdgeDeprecation, EdgeID: 2, Doc: doc, Version: 2},
		{Codec: EdgeTombstone, EdgeID: 3, Doc: doc, Version: 2, Reason: 4},
	} {
		got, err := DecodePublication(EncodePublication(p))
		if err != nil || !reflect.DeepEqual(got, p) {
			t.Fatalf("%+v: %+v %v", p, got, err)
		}
	}
}
