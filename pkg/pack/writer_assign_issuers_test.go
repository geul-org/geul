//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterAssignIssuers — new Issuers are appended in Dir order; known ones keep their number
package pack

import (
	"reflect"
	"testing"
)

func TestWriterAssignIssuers(t *testing.T) {
	w := &writer{issuers: []string{"0x00010000"}, inum: map[string]uint32{"0x00010000": 0}, puts: []*put{
		{ref: pathRef{space: spaceP, issuer: "0x0001"}}, {ref: pathRef{space: spaceP, issuer: "0x0000"}},
		{ref: pathRef{space: spaceP, issuer: "0x0001"}}, {ref: pathRef{space: spaceP, issuer: "0x00010000"}}, {ref: pathRef{}},
	}}
	w.assignIssuers()
	if !reflect.DeepEqual(w.issuers, []string{"0x00010000", "0x0000", "0x0001"}) || w.inum["0x0001"] != 2 {
		t.Fatal(w.issuers, w.inum)
	}
}
