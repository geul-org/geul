//ff:func feature=notation type=decoder control=selection
//ff:what formatTime — GREGORIAN-DATE(YYYY-MM-DD) or POSIX-TIMESTAMP-SECONDS(n)
package notation

import (
	"fmt"
	"strconv"

	"github.com/geul-org/geul/go/pkg/core"
)

// formatTime writes a GREGORIAN-DATE as its proleptic Gregorian day (year 0
// is 1 BC, a negative year carries `-`) and a POSIX-TIMESTAMP-SECONDS as its
// seconds.
func formatTime(r core.Record) (string, bool) {
	t, err := core.DecodeTimeValue(r)
	switch {
	case err != nil:
		return "", false
	case r.Codec == core.CodecPosixSeconds:
		return "POSIX-TIMESTAMP-SECONDS(" + strconv.FormatInt(t.Unix(), 10) + ")", true
	}
	y := t.Year()
	sign := ""
	if y < 0 {
		sign, y = "-", -y
	}
	return fmt.Sprintf("GREGORIAN-DATE(%s%04d-%02d-%02d)", sign, y, int(t.Month()), t.Day()), true
}
