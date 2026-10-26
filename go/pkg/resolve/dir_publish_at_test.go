//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirPublishAt — versions at an address with any document ID; a republish returns the entry, the next chains; exhausted, broken-list and unwritable cases fail
package resolve

import (
	"github.com/geul-org/geul/go/pkg/core"
	"math"
	"testing"
)

func TestDirPublishAt(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	k := codecKey(0x1018)
	doc := codecDocID(0x1018)
	b := codecDocOf(0x1018, 4)
	ws, _ := core.ReadWords(b)
	h, _ := core.ContentHash(ws)
	if e, err := d.publishAt(k, []core.Word{0}, doc, b, h); err != nil || e.Version != 1 {
		t.Fatalf("first: %+v %v", e, err)
	}
	if e, err := d.publishAt(k, []core.Word{0}, doc, b, h); err != nil || e.Version != 1 {
		t.Fatalf("again: %+v %v", e, err)
	}
	if lg, err := d.readList(k, core.FormatID(doc)); err != nil || len(lg.Entries) != 1 {
		t.Fatalf("list: %+v %v", lg, err)
	}
	b2 := codecDocOf(0x1018, 4, 2)
	ws2, _ := core.ReadWords(b2)
	h2, _ := core.ContentHash(ws2)
	if e, err := d.publishAt(k, []core.Word{0}, doc, b2, h2); err != nil || e.Version != 2 {
		t.Fatalf("second: %+v %v", e, err)
	}
	full := codecKey(0x101A)
	d.write(listPath(full), encodeVersionList([]core.Record{core.EncodePublication(core.Publication{Codec: core.EdgePublication, EdgeID: 1, Issuer: []core.Word{0}, Doc: doc, Version: math.MaxUint32})}, nil, 0, 0))
	if _, err := d.publishAt(full, []core.Word{0}, doc, b, h); err == nil {
		t.Fatal("exhausted: want error")
	}
	junk := codecKey(0x101B)
	d.write(listPath(junk), []byte("x"))
	if _, err := d.publishAt(junk, []core.Word{0}, doc, b, h); err == nil {
		t.Fatal("bad list: want error")
	}
	blockWrite(t, d, listPath(codecKey(0x1019)))
	if _, err := d.publishAt(codecKey(0x1019), []core.Word{0}, doc, b, h); err == nil {
		t.Fatal("blocked: want error")
	}
}
