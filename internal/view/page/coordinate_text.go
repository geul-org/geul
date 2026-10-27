//ff:func feature=page type=helper control=selection
//ff:what coordinateText — a coordinate as "lat, lon ±u unit" with its body first when it is not Earth
package page

import (
	"strconv"

	"github.com/geul-org/gogeul/pkg/inspect"
)

// coordinateText shows a COORDINATE-EARTH or COORDINATE-GEO: the body (Qn or
// codec and hex) first for COORDINATE-GEO, then latitude and longitude in
// full binary64 precision and the uncertainty in shortest binary32 form in its unit (Qn or codec and
// hex; ±? when the unit is ABSENT, the uncertainty unknown).
func coordinateText(v inspect.Value) string {
	s := strconv.FormatFloat(v.Lat, 'g', -1, 64) + ", " + strconv.FormatFloat(v.Lon, 'g', -1, 64)
	switch {
	case v.Unit == "":
		s += " ±?"
	default:
		s += " ±" + strconv.FormatFloat(v.Float, 'g', -1, 32) + " " + v.Unit
	}
	if v.Globe != "" {
		return v.Globe + " " + s
	}
	return s
}
