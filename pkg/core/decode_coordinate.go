//ff:func feature=core type=decoder control=selection
//ff:what DecodeCoordinate — the globe, latitude, longitude and uncertainty of a COORDINATE-EARTH or COORDINATE-GEO Value
package core

import (
	"errors"
	"math"
)

// DecodeCoordinate reads a COORDINATE-EARTH (Payload exactly 10 WORDs) or
// COORDINATE-GEO (a whole Reference record, then 10 WORDs) Value. Ranges and
// finiteness are not checked here (schema V8).
func DecodeCoordinate(r Record) (Coordinate, error) {
	var c Coordinate
	p := r.Payload
	switch {
	case r.Kind != KindValue || (r.Codec != CodecCoordinateEarth && r.Codec != CodecCoordinateGeo):
		return Coordinate{}, errors.New("geul/core: not a COORDINATE Value")
	case r.Codec == CodecCoordinateGeo:
		g, n, err := DecodeRecord(p)
		if err != nil || g.Kind != KindReference {
			return Coordinate{}, errors.New("geul/core: COORDINATE-GEO without a globe Reference")
		}
		c.Globe, p = &g, p[n:]
	}
	if len(p) != 10 {
		return Coordinate{}, errors.New("geul/core: COORDINATE needs 10 WORDs after the globe")
	}
	f := func(w []Word) float64 {
		return math.Float64frombits(uint64(w[0])<<48 | uint64(w[1])<<32 | uint64(w[2])<<16 | uint64(w[3]))
	}
	c.Lat, c.Lon, c.U = f(p[0:4]), f(p[4:8]), uint32(p[8])<<16|uint32(p[9])
	return c, nil
}
