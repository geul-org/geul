//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestParseIssuers — bad count, bad width, truncated ID and duplicates fail
package pack

import (
	"errors"
	"testing"
)

func TestParseIssuers(t *testing.T) {
	bad := [][]byte{{9, 0, 0, 0}, {1, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0}, {1, 0, 0, 0, 2, 0, 0}, {2, 0, 0, 0, 1, 0, 0, 1, 0, 0}, {1}}
	for i, b := range bad {
		if _, err := parseIssuers(&cursor{b: b}); !errors.Is(err, ErrCorrupt) {
			t.Fatal(i, err)
		}
	}
}
