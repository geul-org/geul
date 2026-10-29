//ff:func feature=pack type=helper control=sequence
//ff:what TestDirFileClose — nil
package pack

import "testing"

func TestDirFileClose(t *testing.T) {
	if (&dirFile{}).Close() != nil {
		t.Fatal("close")
	}
}
