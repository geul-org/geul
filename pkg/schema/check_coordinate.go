//ff:func feature=schema type=rule control=selection
//ff:what checkCoordinate — V8: a COORDINATE-EARTH or COORDINATE-GEO has its layout, an Identified globe Reference (Geo), finite degrees in range
package schema

import (
	"fmt"
	"math"

	"github.com/geul-org/gogeul/pkg/core"
)

// checkCoordinate applies V8 to a coordinate Value: COORDINATE-EARTH is
// exactly 10 Payload WORDs; COORDINATE-GEO starts with one whole Reference
// record of Quantifier Identified (the body), then 10 WORDs. Latitude and
// longitude are finite, latitude in [-90, 90], longitude in [-180, 180] on
// Earth and [-180, 360] on another body.
func checkCoordinate(r rec) []Fact {
	c, err := core.DecodeCoordinate(r.R)
	lonMax := 180.0
	if r.R.Codec == core.CodecCoordinateGeo {
		lonMax = 360
	}
	switch {
	case err != nil:
		return []Fact{factAt("V8", Error, r, "Payload", "Globe Reference(Geo)와 10 WORD", err.Error())}
	case c.Globe != nil && c.Globe.Codec&7 != 0:
		return []Fact{factAt("V8", Error, r, "Payload", "Globe Quantifier Identified", fmt.Sprintf("0x%02X", c.Globe.Codec))}
	case math.IsNaN(c.Lat) || math.IsNaN(c.Lon) || math.IsInf(c.Lat, 0) || math.IsInf(c.Lon, 0):
		return []Fact{factAt("V8", Error, r, "Payload", "유한한 위도·경도", fmt.Sprintf("%v, %v", c.Lat, c.Lon))}
	case c.Lat < -90 || c.Lat > 90 || c.Lon < -180 || c.Lon > lonMax:
		return []Fact{factAt("V8", Error, r, "Payload", fmt.Sprintf("위도 [-90, 90], 경도 [-180, %v]", lonMax), fmt.Sprintf("%v, %v", c.Lat, c.Lon))}
	default:
		return nil
	}
}
