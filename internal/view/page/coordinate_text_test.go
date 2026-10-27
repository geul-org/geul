//ff:func feature=page type=helper control=sequence
//ff:what TestCoordinateText — Earth with metres, another body first, unknown uncertainty as ±?
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestCoordinateText(t *testing.T) {
	if got := coordinateText(inspect.Value{Codec: core.CodecCoordinateEarth, Lat: 52.516666666667, Lon: 13.383333333333, U: 2172487}); got != "52.516666666667, 13.383333333333 ±2172.487m" {
		t.Fatalf("earth: %q", got)
	}
	if got := coordinateText(inspect.Value{Codec: core.CodecCoordinateGeo, Lat: 4.5, Lon: 137.4, U: core.UncertaintyUnknown, Globe: "Q111"}); got != "Q111 4.5, 137.4 ±?" {
		t.Fatalf("geo: %q", got)
	}
}
