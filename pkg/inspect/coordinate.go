//ff:func feature=inspect type=decoder control=sequence
//ff:what coordinate — the latitude, longitude, uncertainty, its unit and the body of a COORDINATE-EARTH or COORDINATE-GEO
package inspect

import "github.com/geul-org/gogeul/pkg/core"

// coordinate decodes a coordinate Value (core.DecodeCoordinate) into a Value
// with latitude, longitude, the uncertainty (Float), its unit and the body:
// a Reference as referenceText (a WIKIDATA-QID as Qn, any other Reference as
// its codec name and payload hex); the body is "" for COORDINATE-EARTH and
// the unit "" when it is ABSENT (uncertainty unknown).
func coordinate(r core.Record) (Value, error) {
	c, err := core.DecodeCoordinate(r)
	if err != nil {
		return Value{}, err
	}
	v := Value{Codec: r.Codec, Lat: c.Lat, Lon: c.Lon, Float: float64(c.Uncertainty)}
	if c.Globe != nil {
		v.Globe = referenceText(*c.Globe)
	}
	if c.Unit != nil {
		v.Unit = referenceText(*c.Unit)
	}
	return v, nil
}
