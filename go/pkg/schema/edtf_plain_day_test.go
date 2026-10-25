//ff:func feature=schema type=helper control=iteration dimension=1
//ff:what TestEdtfPlainDay — plain and BCE days match; reduced precision, qualifiers, unspecified digits, intervals and odd shapes do not
package schema

import "testing"

func TestEdtfPlainDay(t *testing.T) {
	for s, want := range map[string]bool{
		"1952-03-11": true, "-0043-03-15": true,
		"1952-03": false, "1952": false, "1952-03-11?": false, "19XX-03-11": false, "1952-03-11/1953": false,
		"1952/03/11": false, "1952-03-1a": false, "": false,
	} {
		if edtfPlainDay([]byte(s)) != want {
			t.Fatalf("%q: want %v", s, want)
		}
	}
}
