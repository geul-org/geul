//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what TestParseRawRecord — exactly one record; bad hex, a broken record and extra words fail
package notation

import "testing"

func TestParseRawRecord(t *testing.T) {
	if r, err := parseRawRecord("0x00880003"); err != nil || r.Payload[0] != 3 {
		t.Fatalf("%+v %v", r, err)
	}
	for _, s := range []string{"0x1", "0x0b86", "0x0088000300880003"} {
		if _, err := parseRawRecord(s); err == nil {
			t.Errorf("%s: want error", s)
		}
	}
}
