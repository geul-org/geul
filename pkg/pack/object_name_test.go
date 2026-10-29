//ff:func feature=pack type=helper control=sequence
//ff:what TestObjectName — o/<hex>.gpk for data packs, .gpx for index objects
package pack

import (
	"strings"
	"testing"
)

func TestObjectName(t *testing.T) {
	if n := objectName(object{kind: kindData}); n != "o/"+strings.Repeat("0", 64)+".gpk" {
		t.Fatal(n)
	}
	if n := objectName(object{kind: kindIndex, hash: [32]byte{0xab}}); n != "o/ab"+strings.Repeat("0", 62)+".gpx" {
		t.Fatal(n)
	}
}
