//ff:func feature=notation type=encoder control=sequence
//ff:what TestLookupOf — the context's Lookup when set; otherwise IDs only
package notation

import "testing"

func TestLookupOf(t *testing.T) {
	if _, err := lookupOf(Context{Lookup: fakeLookup})("geul.org/quantity"); err != nil {
		t.Fatal(err)
	}
	if _, err := lookupOf(Context{})("geul.org/quantity"); err == nil {
		t.Fatal("names without a Lookup: want error")
	}
}
