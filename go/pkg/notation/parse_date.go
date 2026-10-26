//ff:func feature=notation type=encoder control=sequence
//ff:what parseDate — a GREGORIAN-DATE Value from YYYY-MM-DD (proleptic Gregorian, year 0 = 1 BC, - for earlier)
package notation

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/geul-org/geul/go/pkg/core"
)

// datePattern is an optional sign, at least 4 year digits, month and day.
var datePattern = regexp.MustCompile(`^(-?[0-9]{4,})-([0-9]{2})-([0-9]{2})$`)

// parseDate encodes the day as GREGORIAN-DATE: days from 1970-01-01 in a
// signed 32-bit integer over two words. A day that does not exist or is out
// of range is an error.
func parseDate(s string) (core.Record, error) {
	m := datePattern.FindStringSubmatch(s)
	if m == nil {
		return core.Record{}, fmt.Errorf("geul/notation: %q is not YYYY-MM-DD", s)
	}
	y, _ := strconv.Atoi(m[1])
	mo, _ := strconv.Atoi(m[2])
	d, _ := strconv.Atoi(m[3])
	t := time.Date(y, time.Month(mo), d, 0, 0, 0, 0, time.UTC)
	days := t.Unix() / 86400
	if t.Year() != y || int(t.Month()) != mo || t.Day() != d || days < -1<<31 || days > 1<<31-1 {
		return core.Record{}, fmt.Errorf("geul/notation: %q is not a day GREGORIAN-DATE holds", s)
	}
	u := uint32(int32(days))
	return core.Record{Kind: core.KindValue, Codec: core.CodecGregorianDate, Payload: []core.Word{core.Word(u >> 16), core.Word(u)}}, nil
}
