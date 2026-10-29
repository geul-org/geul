//ff:func feature=pack type=helper control=sequence
//ff:what TestDictBytes — nil for no dictionary, the bytes otherwise
package pack

import "testing"

func TestDictBytes(t *testing.T) {
	if dictBytes(nil) != nil || string(dictBytes(&Dict{Bytes: []byte("d")})) != "d" {
		t.Fatal("dictBytes")
	}
}
