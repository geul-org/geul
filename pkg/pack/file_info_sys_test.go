//ff:func feature=pack type=helper control=sequence
//ff:what TestFileInfoSys — nil
package pack

import "testing"

func TestFileInfoSys(t *testing.T) {
	if (fileInfo{}).Sys() != nil {
		t.Fatal("sys")
	}
}
