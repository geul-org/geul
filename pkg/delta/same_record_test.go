//ff:func feature=delta type=helper control=sequence
//ff:what TestSameRecord — equal encodings are the same; different payloads are not
package delta

import "testing"

func TestSameRecord(t *testing.T) {
	if !sameRecord(bindOf(1, 2, "a"), bindOf(1, 2, "a")) || sameRecord(bindOf(1, 2, "a"), bindOf(1, 2, "b")) {
		t.Fatal("compare")
	}
}
