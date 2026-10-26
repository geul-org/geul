//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what TestParseQuantity — four parts; wrong arity, an unreadable part and a part of the wrong kind fail
package notation

import "testing"

func TestParseQuantity(t *testing.T) {
	if _, err := parseQuantity([]string{"1", "Q1", "0", "2"}, Context{}); err != nil {
		t.Fatal(err)
	}
	for _, args := range [][]string{{"1"}, {"1", "word", "_", "_"}, {`"a"`, "_", "_", "_"}} {
		if _, err := parseQuantity(args, Context{}); err == nil {
			t.Errorf("%v: want error", args)
		}
	}
}
