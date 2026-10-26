//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what TestParseValueArg — each codec's argument; codecs without a call form fail
package notation

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestParseValueArg(t *testing.T) {
	cases := map[uint8]string{
		core.CodecLanguageBCP47Even: `"en"`, core.CodecUnsignedDecimal: "1.5", core.CodecUnsignedInteger: "9",
		core.CodecGregorianDate: "2026-10-27", core.CodecPosixSeconds: "100", core.CodecSHA256Digest: "0x" + "00000000000000000000000000000000000000000000000000000000000000ab",
	}
	for c, arg := range cases {
		if _, err := parseValueArg(c, arg); err != nil {
			t.Errorf("0x%02X %s: %v", c, arg, err)
		}
	}
	if _, err := parseValueArg(core.CodecBinaryFloat, "1"); err == nil {
		t.Fatal("no call form: want error")
	}
}
