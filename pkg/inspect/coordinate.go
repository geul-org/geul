//ff:func feature=inspect type=decoder control=sequence
//ff:what coordinate — the latitude, longitude, uncertainty and body of a COORDINATE-EARTH or COORDINATE-GEO
package inspect

import "github.com/geul-org/gogeul/pkg/core"

// coordinate decodes a coordinate Value (core.DecodeCoordinate) into
// latitude, longitude, uncertainty and its body: "" for COORDINATE-EARTH, a
// WIKIDATA-QID globe as Qn, any other Reference as its codec name and payload
// hex.
func coordinate(r core.Record) (float64, float64, uint32, string, error) {
	c, err := core.DecodeCoordinate(r)
	if err != nil {
		return 0, 0, 0, "", err
	}
	globe := ""
	if c.Globe != nil {
		globe = referenceText(*c.Globe)
	}
	return c.Lat, c.Lon, c.U, globe, nil
}
