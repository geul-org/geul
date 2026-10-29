//ff:func feature=pack type=helper control=sequence
//ff:what TestDirPrefix — "" for the top, a trailing slash otherwise
package pack

import "testing"

func TestDirPrefix(t *testing.T) {
	if dirPrefix(".") != "" || dirPrefix("a/b") != "a/b/" {
		t.Fatal("dirPrefix")
	}
}
