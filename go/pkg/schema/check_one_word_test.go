//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckOneWord — 1 WORD passes, longer fails D4
package schema

import (
	"testing"
)

func TestCheckOneWord(t *testing.T) {
	if checkOneWord(rec{Total: 1}, "END") != nil || rulesOf(checkOneWord(rec{Total: 2}, "END")) != "D4" {
		t.Fatal("wrong")
	}
}
