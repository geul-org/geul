//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what TestOpenDirResolve — a root importing codec spec documents of the published pack archive (geul-org/artifacts/geul) resolves through a GEULPROXY file:// source read by OpenDir
package pack

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/resolve"
)

// TestOpenDirResolve reads the codec spec documents published at geul-dev
// draft/geul-org/artifacts/geul/.well-known/geul/ (pack sets, the stand-in for
// geul.org/.well-known/geul/): a root document imports the PUBLICATION spec
// document (latest) and pins the FIELD one at version 1. It is skipped when
// this module is checked out on its own, without geul-dev draft/ around it.
func TestOpenDirResolve(t *testing.T) {
	ctx := context.Background()
	root, _ := filepath.Abs("../../../geul-org/artifacts/geul")
	wk := filepath.Join(root, ".well-known", "geul")
	if !IsPacked(wk) {
		t.Skip("geul-org/artifacts/geul/.well-known/geul holds no pack sets")
	}
	c, err := resolve.ConfigFromEnv(func(k string) string {
		return map[string]string{"GEULPROXY": "file://" + root + ",off"}[k]
	})
	a, aerr := OpenDir(ctx, wk, BuildOptions{})
	if err != nil || aerr != nil {
		t.Fatal(err, aerr)
	}
	c.Proxy[0].Source = resolve.Dir{Root: root, Store: a}
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	id, _ := core.EncodeDocID([]core.Word{0}, []core.Word{0xffff, 0, 0, 7})
	in := enc(id)
	for i, imp := range [][2]core.Word{{0x1018, 0}, {0x101C, 1}} {
		ptr, _ := core.EncodeDocPointer([]core.Word{0}, []core.Word{0, 0, 0, imp[0]}, []core.Word{0, imp[1]})
		in = append(in, enc(core.Record{Kind: core.KindMeta, Codec: core.MetaImportDoc, Payload: append(enc(ptr), core.Word(i+1))})...)
	}
	in = append(append(in, 0x180E), enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("root")))...)
	g, err := (&resolve.Resolver{Config: c}).Resolve(ctx, core.PutWords(core.WrapDocument(append(in, 0x180F))))
	if err != nil {
		t.Fatal(err)
	}
	key := func(w core.Word) resolve.NodeKey {
		return resolve.NodeKey{DocKey: resolve.DocKey{Issuer: core.FormatID([]core.Word{0}), Doc: core.FormatID([]core.Word{0, 0, 0, w})}, Version: 1}
	}
	pub, field := g.Docs[key(0x1018)], g.Docs[key(0x101C)]
	if len(g.Docs) != 2 || pub == nil || field == nil || pub.Source != "file://"+root {
		t.Fatalf("got %+v", g.Docs)
	}
}
