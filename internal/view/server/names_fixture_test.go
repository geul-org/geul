//ff:func feature=server type=helper control=sequence
//ff:what namesFixture — test fixture: a tree with issuers.geul naming 0x0000 "geul.org" and documents.geul naming its document 0xffff000000000002 "sample-b", and its codec spec document 0x0000000000001018 "publication"
package server

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func namesFixture(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	top := filepath.Join(root, ".well-known", "geul")
	doc := filepath.Join(top, "0x0000", "0xffff000000000002")
	if err := os.MkdirAll(doc, 0o755); err != nil {
		t.Fatal(err)
	}
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	subj := core.Record{Kind: core.KindReference, Codec: core.RefIssuerPointer << 3, Payload: []core.Word{0}}
	p := append(enc(core.EncodeLocalRef(1)), enc(subj)...)
	p = append(p, enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("geul.org")))...)
	in := append([]core.Word{0x180E}, enc(core.Record{Kind: core.KindEdge, Codec: core.EdgeBind, Payload: p})...)
	os.WriteFile(filepath.Join(top, "issuers.geul"), core.PutWords(core.WrapDocument(append(in, 0x180F))), 0o644)
	os.WriteFile(filepath.Join(doc, "0x00000001.geul"), core.PutWords(core.WrapDocument([]core.Word{0x180F})), 0o644)
	dsubj := core.Record{Kind: core.KindReference, Codec: core.RefDocPointer << 3, Payload: []core.Word{0, 0xffff, 0, 0, 2, 0, 0}}
	dp := append(enc(core.EncodeLocalRef(1)), enc(dsubj)...)
	dp = append(dp, enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("sample-b")))...)
	csubj := core.Record{Kind: core.KindReference, Codec: core.RefDocPointer << 3, Payload: []core.Word{0, 0, 0, 0, 0x1018, 0, 0}}
	cp := append(enc(core.EncodeLocalRef(2)), enc(csubj)...)
	cp = append(cp, enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("publication")))...)
	din := append([]core.Word{0x180E}, enc(core.Record{Kind: core.KindEdge, Codec: core.EdgeBind, Payload: dp})...)
	din = append(din, enc(core.Record{Kind: core.KindEdge, Codec: core.EdgeBind, Payload: cp})...)
	os.WriteFile(filepath.Join(top, "0x0000", "documents.geul"), core.PutWords(core.WrapDocument(append(din, 0x180F))), 0o644)
	os.MkdirAll(filepath.Join(top, "0x0000", "0x0000000000001018"), 0o755)
	return root
}
