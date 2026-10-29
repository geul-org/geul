//ff:func feature=pack type=encoder control=sequence
//ff:what TestEncodeInventory — round trip through parseInventory
package pack

import (
	"reflect"
	"testing"
)

func TestEncodeInventory(t *testing.T) {
	inv := sampleInventory()
	got, err := parseInventory(encodeInventory(inv))
	if err != nil || !reflect.DeepEqual(got, inv) {
		t.Fatal(got, err)
	}
}
