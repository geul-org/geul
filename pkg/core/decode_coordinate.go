//ff:func feature=core type=decoder control=selection
//ff:what DecodeCoordinate — the globe, latitude, longitude, uncertainty and unit of a COORDINATE-EARTH or COORDINATE-GEO Value
package core

import (
	"errors"
	"math"
)

// DecodeCoordinate reads a COORDINATE-EARTH (10 WORDs, then the unit) or
// COORDINATE-GEO (a whole globe Reference record, 10 WORDs, then the unit)
// Value: latitude and longitude binary64, the uncertainty binary32. The unit is one whole Reference
// record ending the payload, or the slot marker ABSENT (Unit nil: uncertainty
// unknown). Ranges, finiteness, the uncertainty's sign and the +0 an unknown
// uncertainty needs are not checked here (schema V8).
func DecodeCoordinate(r Record) (Coordinate, error) {
	var c Coordinate
	p := r.Payload
	switch {
	case r.Kind != KindValue || (r.Codec != CodecCoordinateEarth && r.Codec != CodecCoordinateGeo):
		return Coordinate{}, errors.New("geul/core: not a COORDINATE Value")
	case r.Codec == CodecCoordinateGeo:
		g, n, err := DecodeRecord(p)
		if err != nil || g.Kind != KindReference || IsSlotMarker(g) {
			return Coordinate{}, errors.New("geul/core: COORDINATE-GEO without a globe Reference")
		}
		c.Globe, p = &g, p[n:]
	}
	if len(p) < 11 {
		return Coordinate{}, errors.New("geul/core: COORDINATE needs 10 WORDs and a unit after the globe")
	}
	u, n, err := DecodeRecord(p[10:])
	if err != nil || u.Kind != KindReference || 10+n != len(p) || (IsSlotMarker(u) && !IsAbsent(u)) {
		return Coordinate{}, errors.New("geul/core: COORDINATE unit is not one Reference or ABSENT ending the payload")
	}
	if !IsAbsent(u) {
		c.Unit = &u
	}
	f := func(w []Word) float64 {
		return math.Float64frombits(uint64(w[0])<<48 | uint64(w[1])<<32 | uint64(w[2])<<16 | uint64(w[3]))
	}
	c.Lat, c.Lon = f(p[0:4]), f(p[4:8])
	c.Uncertainty = math.Float32frombits(uint32(p[8])<<16 | uint32(p[9]))
	return c, nil
}
