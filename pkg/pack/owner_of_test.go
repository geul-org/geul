//ff:func feature=pack type=writer control=sequence
//ff:what TestOwnerOf — the Issuer of Issuer paths, "" for top-level and h/ paths
package pack

import "testing"

func TestOwnerOf(t *testing.T) {
	if ownerOf("0x0000/documents.geul") != "0x0000" || ownerOf("issuers.geul") != "" || ownerOf("h/x.geul") != "" || ownerOf("0x0000") != "" {
		t.Fatal("ownerOf")
	}
}
