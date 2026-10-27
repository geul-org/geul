//ff:func feature=inspect type=decoder control=sequence
//ff:what TestDecode — a full document interprets HEAD and BODY; odd bytes and a non-document fail
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestDecode(t *testing.T) {
	d, err := Decode(sampleDoc(t))
	if err != nil {
		t.Fatal(err)
	}
	if len(d.Head) != 6 || len(d.Body) != 8 || d.Profile != "tangeul-1.0" || !d.Toulmin {
		t.Fatalf("got %+v", d)
	}
	if _, ok := d.Body[3].Data.(toulmin.NodePacket); !ok || d.Body[3].Name != "CASE-DECLARE" || d.Body[4].ID != 12 {
		t.Fatalf("body: %+v", d.Body[3:5])
	}
	if _, err := Decode([]byte{1}); err == nil {
		t.Fatal("odd bytes: want error")
	}
	if _, err := Decode([]byte{0x18, 0x0F}); err == nil {
		t.Fatal("not a document: want error")
	}
}
