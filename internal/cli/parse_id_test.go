//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestParseID — #n; no #, not a number and #0 fail
package cli

import "testing"

func TestParseID(t *testing.T) {
	if n, err := parseID("#12"); err != nil || n != 12 {
		t.Fatal(err)
	}
	for _, s := range []string{"12", "#x", "#0"} {
		if _, err := parseID(s); err == nil {
			t.Errorf("%s: want error", s)
		}
	}
}
