//ff:func feature=pack type=encoder control=sequence
//ff:what TestEncodeObjectHeader — 64 bytes: magic, version, kind, size, inventory place, count
package pack

import (
	"reflect"
	"testing"
)

func TestEncodeObjectHeader(t *testing.T) {
	h := objectHeader{kind: kindData, size: 100, invOff: 80, invLen: 20, count: 3}
	b := encodeObjectHeader(h)
	if len(b) != 64 || string(b[:8]) != "GEULPACK" || b[10] != kindData {
		t.Fatalf("%x", b)
	}
	if got, err := parseObjectHeader(b); err != nil || !reflect.DeepEqual(got, h) {
		t.Fatal(got, err)
	}
}
