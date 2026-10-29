//ff:func feature=pack type=helper control=sequence
//ff:what TestMemFileClose — nil
package pack

import "testing"

func TestMemFileClose(t *testing.T) {
	if (&memFile{}).Close() != nil {
		t.Fatal("close")
	}
}
