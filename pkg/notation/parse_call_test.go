//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what TestParseCall — Value calls, QUANTITY, FRACTION and ISSUER-POINTER; wrong arity, unknown names and an unknown Issuer fail
package notation

import "testing"

func TestParseCall(t *testing.T) {
	ctx := Context{Lookup: fakeLookup}
	for _, s := range []string{"UNSIGNED-INTEGER(5)", `EDTF("1971~")`, "ISSUER-POINTER(0x00010000)", "QUANTITY(1, ABSENT, ABSENT, ABSENT)", "FRACTION(2, 3)", "QUANTITY(FRACTION(23.45, 1256.1), ABSENT, ABSENT, ABSENT)"} {
		r, err := parseCall(s, ctx)
		if err != nil || formatRecord(r) != s {
			t.Errorf("%s: %s %v", s, formatRecord(r), err)
		}
	}
	for _, s := range []string{"UNSIGNED-INTEGER(1, 2)", "NOPE(1)", "ISSUER-POINTER(nobody)", "FRACTION(2)"} {
		if _, err := parseCall(s, ctx); err == nil {
			t.Errorf("%s: want error", s)
		}
	}
}
