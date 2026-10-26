//ff:func feature=notation type=encoder control=sequence
//ff:what TestEnumWords — an ENUM name in any case, else raw words
package notation

import "testing"

func TestEnumWords(t *testing.T) {
	enums := map[string]string{"0x0001": "preferred"}
	if ws, err := enumWords("Preferred", enums); err != nil || ws[0] != 1 {
		t.Fatal(err)
	}
	if ws, err := enumWords("0x0003", enums); err != nil || ws[0] != 3 {
		t.Fatal(err)
	}
}
