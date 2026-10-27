//ff:func feature=notation type=encoder control=sequence
//ff:what parseCoordinate — COORDINATE-EARTH from lat, lon, uncertainty, unit, or COORDINATE-GEO from globe and those four; ABSENT unit for an unknown uncertainty
package notation

import (
	"errors"
	"strconv"

	"github.com/geul-org/gogeul/pkg/core"
)

// parseCoordinate reads the arguments of COORDINATE-EARTH (geo false:
// latitude, longitude, uncertainty, unit) or COORDINATE-GEO (geo: the globe
// record first): numbers as binary64, the uncertainty as binary32, the globe
// and unit as records (parseRecord; ABSENT for an unknown uncertainty), and
// checks the result decodes (core.DecodeCoordinate).
func parseCoordinate(args []string, geo bool, ctx Context) (core.Record, error) {
	var c core.Coordinate
	if geo && len(args) == 5 {
		g, err := parseRecord(args[0], ctx)
		if err != nil {
			return core.Record{}, err
		}
		c.Globe, args = &g, args[1:]
	}
	if len(args) != 4 {
		return core.Record{}, errors.New("geul/notation: COORDINATE takes [globe,] latitude, longitude, uncertainty, unit")
	}
	lat, err1 := strconv.ParseFloat(args[0], 64)
	lon, err2 := strconv.ParseFloat(args[1], 64)
	u, err3 := strconv.ParseFloat(args[2], 32)
	unit, err4 := parseRecord(args[3], ctx)
	if err := errors.Join(err1, err2, err3, err4); err != nil {
		return core.Record{}, err
	}
	c.Lat, c.Lon, c.Uncertainty = lat, lon, float32(u)
	if !core.IsAbsent(unit) {
		c.Unit = &unit
	}
	r := core.EncodeCoordinate(c)
	_, err := core.DecodeCoordinate(r)
	return r, err
}
