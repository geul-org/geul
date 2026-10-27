//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what TestParseDecimal — digits kept as C and E for points, exponents and signs; no digits, junk and a huge exponent fail
package notation

import "testing"

func TestParseDecimal(t *testing.T) {
	cases := map[string][2]int64{"1.20": {-2, 120}, "12e2": {2, 12}, "-0.005": {-3, -5}, "+7": {0, 7}, ".5": {-1, 5}, "1.5E-3": {-4, 15}}
	for s, want := range cases {
		e, c, err := parseDecimal(s)
		if err != nil || int64(e) != want[0] || c.Int64() != want[1] {
			t.Errorf("%s: %d %v %v", s, e, c, err)
		}
	}
	for _, s := range []string{"-", "1.2.3", "1e", "1e9999999999"} {
		if _, _, err := parseDecimal(s); err == nil {
			t.Errorf("%s: want error", s)
		}
	}
}
