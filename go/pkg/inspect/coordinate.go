//ff:func feature=inspect type=decoder control=sequence
//ff:what coordinate — the latitude, longitude, uncertainty and body of a COORDINATE-EARTH or COORDINATE-GEO
package inspect

import (
	"strconv"

	"github.com/geul-org/geul/go/pkg/core"
)

// coordinate decodes a coordinate Value (core.DecodeCoordinate) into
// latitude, longitude, uncertainty and its body: "" for COORDINATE-EARTH, a
// WIKIDATA-QID globe as Qn, any other Reference as its codec name and payload
// hex.
func coordinate(r core.Record) (float64, float64, uint32, string, error) {
	c, err := core.DecodeCoordinate(r)
	if err != nil {
		return 0, 0, 0, "", err
	}
	g := c.Globe
	globe := ""
	if g != nil {
		globe = core.CodecName(g.Kind, g.Codec) + " " + core.FormatID(g.Payload)
	}
	if g != nil && g.Codec == core.RefWikidataQID<<3 && len(g.Payload) <= 4 {
		var n uint64
		for _, w := range g.Payload {
			n = n<<16 | uint64(w)
		}
		globe = "Q" + strconv.FormatUint(n, 10)
	}
	return c.Lat, c.Lon, c.U, globe, nil
}
