//ff:func feature=pack type=encoder control=sequence
//ff:what TestAppendHKeys — keys concatenated as they are
package pack

import "testing"

func TestAppendHKeys(t *testing.T) {
	if b := appendHKeys([]byte("x"), []string{"ab", "cd"}); string(b) != "xabcd" {
		t.Fatal(string(b))
	}
}
