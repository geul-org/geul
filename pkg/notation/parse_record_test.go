//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what TestParseRecord — every token form reads back to what formatRecord writes; empty and unreadable tokens fail
package notation

import "testing"

func TestParseRecord(t *testing.T) {
	ctx := Context{Lookup: fakeLookup}
	for _, s := range []string{"_", "UNKNOWN", "NONE", "#3", "Q42", "P31", `"a b"`, `"x"@ko`, "<https://geul.org/>", "1.20", "-3", "0x0000/0x0000000000000806@1", "0x0000/0x0000000000000806#2", "GREGORIAN-DATE(1952-03-11)", "QUANTITY(652, Q828224, _, _)", "0x00900005"} {
		r, err := parseRecord(s, ctx)
		if err != nil || formatRecord(r) != s {
			t.Errorf("%s: %s %v", s, formatRecord(r), err)
		}
	}
	for _, s := range []string{"", "word", "Q", "1.2.3"} {
		if _, err := parseRecord(s, ctx); err == nil {
			t.Errorf("%q: want error", s)
		}
	}
}
