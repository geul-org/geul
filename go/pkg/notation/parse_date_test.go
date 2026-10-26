//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what TestParseDate — days from 1970 both ways; bad forms, days that do not exist and days past 32 bits fail
package notation

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestParseDate(t *testing.T) {
	r, err := parseDate("1970-01-02")
	if err != nil || r.Payload[0] != 0 || r.Payload[1] != 1 {
		t.Fatalf("%+v %v", r, err)
	}
	r, _ = parseDate("1969-12-31")
	if tm, _ := core.DecodeTimeValue(r); tm.Year() != 1969 || tm.Day() != 31 {
		t.Fatal(tm)
	}
	for _, s := range []string{"1970-1-2", "2026-02-30", "99999999-01-01"} {
		if _, err := parseDate(s); err == nil {
			t.Errorf("%s: want error", s)
		}
	}
}
