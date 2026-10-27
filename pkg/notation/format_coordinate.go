//ff:func feature=notation type=decoder control=sequence
//ff:what formatCoordinate — COORDINATE-EARTH(lat, lon, uncertainty, unit) or COORDINATE-GEO(globe, lat, lon, uncertainty, unit) in shortest round-trip numbers, ABSENT for an unknown uncertainty's unit
package notation

import (
	"strconv"
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// formatCoordinate writes a coordinate Value (core.DecodeCoordinate): the
// globe first for COORDINATE-GEO (formatRecord), latitude and longitude as
// the shortest binary64 text, the uncertainty as the shortest binary32 text,
// then the unit (formatRecord; ABSENT when unknown). Each number reads back
// to the same bits (parseCoordinate).
func formatCoordinate(r core.Record) (string, bool) {
	c, err := core.DecodeCoordinate(r)
	if err != nil {
		return "", false
	}
	var parts []string
	if c.Globe != nil {
		parts = append(parts, formatRecord(*c.Globe))
	}
	unit := core.Record{Kind: core.KindReference}
	if c.Unit != nil {
		unit = *c.Unit
	}
	parts = append(parts, strconv.FormatFloat(c.Lat, 'g', -1, 64), strconv.FormatFloat(c.Lon, 'g', -1, 64),
		strconv.FormatFloat(float64(c.Uncertainty), 'g', -1, 32), formatRecord(unit))
	return core.CodecName(core.KindValue, r.Codec) + "(" + strings.Join(parts, ", ") + ")", true
}
