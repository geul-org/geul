//ff:func feature=page type=helper control=sequence
//ff:what TestCoordinateText — Earth with its uncertainty and unit, another body first, unknown uncertainty as ±?
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestCoordinateText(t *testing.T) {
	if got := coordinateText(inspect.Value{Codec: core.CodecCoordinateEarth, Lat: 52.516666666667, Lon: 13.383333333333, Float: 1, Unit: "Q209426"}); got != "52.516666666667, 13.383333333333 ±1 Q209426" {
		t.Fatalf("earth: %q", got)
	}
	if got := coordinateText(inspect.Value{Codec: core.CodecCoordinateGeo, Lat: 4.5, Lon: 137.4, Globe: "Q111"}); got != "Q111 4.5, 137.4 ±?" {
		t.Fatalf("geo: %q", got)
	}
}
