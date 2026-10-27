//ff:func feature=core type=encoder control=iteration dimension=1
//ff:what EncodeCoordinate — a COORDINATE-EARTH Value, or COORDINATE-GEO when a globe Reference is given
package core

import "math"

// EncodeCoordinate builds the coordinate Value of c: COORDINATE-EARTH
// (value/numeric) when c.Globe is nil (WGS 84 Earth), else COORDINATE-GEO with
// the globe Reference first. Latitude and longitude are binary64 (4 WORDs
// each, big-endian), the uncertainty binary32 (2 WORDs), then the unit
// Reference, or the slot marker ABSENT when c.Unit is nil (uncertainty
// unknown). The caller passes finite degrees in range and a finite
// uncertainty ≥ +0 (+0 with no unit); schema checks them (V8).
func EncodeCoordinate(c Coordinate) Record {
	var p []Word
	codec := CodecCoordinateEarth
	if c.Globe != nil {
		p, _ = EncodeRecord(*c.Globe)
		codec = CodecCoordinateGeo
	}
	for _, f := range []float64{c.Lat, c.Lon} {
		b := math.Float64bits(f)
		p = append(p, Word(b>>48), Word(b>>32), Word(b>>16), Word(b))
	}
	u := math.Float32bits(c.Uncertainty)
	p = append(p, Word(u>>16), Word(u))
	unit := []Word{Absent}
	if c.Unit != nil {
		unit, _ = EncodeRecord(*c.Unit)
	}
	return Record{Kind: KindValue, Codec: codec, Payload: append(p, unit...)}
}
