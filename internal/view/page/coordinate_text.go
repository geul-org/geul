//ff:func feature=page type=helper control=selection
//ff:what coordinateText — a coordinate as "lat, lon ±u" with its body first when it is not Earth
package page

import (
	"strconv"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// coordinateText shows a COORDINATE-EARTH or COORDINATE-GEO: the body (Qn or
// codec and hex) first for COORDINATE-GEO, then latitude and longitude in
// full binary64 precision and the uncertainty in metres (±? when unknown).
func coordinateText(v inspect.Value) string {
	s := strconv.FormatFloat(v.Lat, 'g', -1, 64) + ", " + strconv.FormatFloat(v.Lon, 'g', -1, 64)
	switch {
	case v.U == core.UncertaintyUnknown:
		s += " ±?"
	default:
		s += " ±" + strconv.FormatFloat(float64(v.U)/1000, 'f', -1, 64) + "m"
	}
	if v.Globe != "" {
		return v.Globe + " " + s
	}
	return s
}
