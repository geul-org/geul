//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirPublish — sequential versions, idempotent republish, list and hash files written; bad, unnamed, exhausted and unwritable inputs fail
package resolve

import (
	"context"
	"github.com/geul-org/geul/go/pkg/core"
	"math"
	"testing"
)

func TestDirPublish(t *testing.T) {
	d := sampleTree(t)
	k := sampleKey(1, 0).DocKey()
	list, err := d.readList(k)
	if err != nil || len(list) != 2 || list[0].Version != 1 || list[1].Version != 2 {
		t.Fatalf("list: %+v %v", list, err)
	}
	if e, err := d.Publish(sampleDoc(1, "A v1")); err != nil || e.Version != 1 {
		t.Fatalf("republish: %+v %v", e, err)
	}
	if b, err := d.Fetch(context.Background(), hashPath(list[1].Hash)); err != nil || len(b) == 0 {
		t.Fatalf("hash path: %v", err)
	}
	if _, err := d.Publish([]byte{1}); err == nil {
		t.Fatal("odd bytes: want error")
	}
	if _, err := d.Publish(core.PutWords(core.WrapDocument([]core.Word{0x180E}))); err == nil {
		t.Fatal("no END: want error")
	}
	if _, err := d.Publish(core.PutWords(core.WrapDocument([]core.Word{0x180E, 0x180F}))); err == nil {
		t.Fatal("unnamed: want error")
	}
	full := sampleKey(4, 0).DocKey()
	if err := d.write(listPath(full), []byte(formatVersionList([]VersionEntry{{Version: math.MaxUint32}}))); err != nil {
		t.Fatal(err)
	}
	if _, err := d.Publish(sampleDoc(4, "x")); err == nil {
		t.Fatal("exhausted: want error")
	}
	if err := d.write(listPath(sampleKey(5, 0).DocKey()), []byte("junk")); err != nil {
		t.Fatal(err)
	}
	if _, err := d.Publish(sampleDoc(5, "x")); err == nil {
		t.Fatal("bad list: want error")
	}
}
