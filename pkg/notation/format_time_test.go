//ff:func feature=notation type=decoder control=sequence
//ff:what TestFormatTime — dates (BC years signed), seconds; a broken payload is false
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFormatTime(t *testing.T) {
	d, _ := parseDate("1952-03-11")
	if s, ok := formatTime(d); !ok || s != "GREGORIAN-DATE(1952-03-11)" {
		t.Fatal(s)
	}
	bc, _ := parseDate("-0044-03-15")
	if s, _ := formatTime(bc); s != "GREGORIAN-DATE(-0044-03-15)" {
		t.Fatal(s)
	}
	if s, _ := formatTime(core.EncodePosixSeconds(-5)); s != "POSIX-TIMESTAMP-SECONDS(-5)" {
		t.Fatal(s)
	}
	if _, ok := formatTime(core.Record{Kind: core.KindValue, Codec: core.CodecGregorianDate, Payload: []core.Word{1}}); ok {
		t.Fatal("broken: want false")
	}
}
