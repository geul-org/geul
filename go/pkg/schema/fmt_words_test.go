//ff:func feature=schema type=rule control=sequence
//ff:what TestFmtWords — N WORD
package schema

import (
	"testing"
)

func TestFmtWords(t *testing.T) {
	if fmtWords(3) != "3 WORD" {
		t.Fatal(fmtWords(3))
	}
}
