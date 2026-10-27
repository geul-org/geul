//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCheckCoordinate — valid Earth and Geo points pass; bad width, no or quantified globe, non-finite and out-of-range degrees fail
package schema

import (
	"math"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCheckCoordinate(t *testing.T) {
	q := func(n core.Word, quant uint8) core.Record {
		return core.Record{Kind: core.KindReference, Codec: core.RefWikidataQID<<3 | quant, Payload: []core.Word{n}}
	}
	geo := func(g core.Record, lat, lon float64) core.Record {
		return core.EncodeCoordinate(core.Coordinate{Globe: &g, Lat: lat, Lon: lon})
	}
	cases := []struct {
		r    core.Record
		want string
	}{
		{core.EncodeCoordinate(core.Coordinate{Lat: 52.5, Lon: 13.4, U: 44}), ""},
		{geo(q(111, 0), 4.5, 359), ""},
		{core.Record{Kind: core.KindValue, Codec: 0x04, Payload: make([]core.Word, 9)}, "V8"},
		{geo(q(111, 1), 4.5, 10), "V8"},
		{core.EncodeCoordinate(core.Coordinate{Lat: math.NaN(), Lon: 1}), "V8"},
		{core.EncodeCoordinate(core.Coordinate{Lat: 1, Lon: math.Inf(1)}), "V8"},
		{core.EncodeCoordinate(core.Coordinate{Lat: 91, Lon: 1}), "V8"},
		{core.EncodeCoordinate(core.Coordinate{Lat: 1, Lon: 200}), "V8"},
		{geo(q(111, 0), 1, 361), "V8"},
	}
	for i, c := range cases {
		if got := rulesOf(checkCoordinate(recOf(c.r))); got != c.want {
			t.Fatalf("case %d: %q want %q", i, got, c.want)
		}
	}
}
