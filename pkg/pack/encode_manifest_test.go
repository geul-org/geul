//ff:func feature=pack type=encoder control=sequence
//ff:what TestEncodeManifest — magic, version, six sections; parseManifest gives the same manifest back
package pack

import (
	"reflect"
	"testing"
)

func TestEncodeManifest(t *testing.T) {
	m := sampleManifest()
	b := encodeManifest(m)
	if string(b[:8]) != "GEULMANI" || b[8] != 1 || b[12] != 6 {
		t.Fatalf("%x", b[:16])
	}
	got, err := parseManifest(b)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, m) {
		t.Fatalf("%+v\n%+v", got, m)
	}
}
