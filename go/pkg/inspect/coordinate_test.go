//ff:func feature=inspect type=decoder control=sequence
//ff:what TestCoordinate — Earth has no body, a QID body reads as Qn, another Reference by codec and hex; a broken value fails
package inspect

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestCoordinate(t *testing.T) {
	lat, lon, u, g, err := coordinate(core.EncodeCoordinate(core.Coordinate{Lat: 52.5, Lon: 13.4, U: 44}))
	if err != nil || lat != 52.5 || lon != 13.4 || u != 44 || g != "" {
		t.Fatalf("earth: %v %v %v %q %v", lat, lon, u, g, err)
	}
	mars := core.Record{Kind: core.KindReference, Codec: core.RefWikidataQID << 3, Payload: []core.Word{111}}
	if _, _, _, g, err := coordinate(core.EncodeCoordinate(core.Coordinate{Globe: &mars, Lat: 1, Lon: 2})); err != nil || g != "Q111" {
		t.Fatalf("mars: %q %v", g, err)
	}
	local := core.EncodeLocalRef(3)
	if _, _, _, g, _ := coordinate(core.EncodeCoordinate(core.Coordinate{Globe: &local, Lat: 1, Lon: 2})); g != "LOCAL 0x0003" {
		t.Fatalf("local: %q", g)
	}
	if _, _, _, _, err := coordinate(core.Record{Kind: core.KindValue, Codec: core.CodecCoordinateEarth}); err == nil {
		t.Fatal("broken: want error")
	}
}
