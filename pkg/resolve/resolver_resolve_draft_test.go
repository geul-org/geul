//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestResolverResolveDraft — a root importing a draft/.well-known/geul codec spec document resolves through a custom GEULPROXY file:// source
package resolve

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// TestResolverResolveDraft reads the codec spec documents published in the GEUL
// spec repository at draft/.well-known/geul/ (the temporary stand-in for
// geul.org/.well-known/geul/): a root document imports the PUBLICATION spec
// document (latest) and pins the FIELD one at version 1. It is skipped when this
// module is checked out on its own, without the spec repository around it.
func TestResolverResolveDraft(t *testing.T) {
	root, _ := filepath.Abs("../../..")
	if _, err := os.Stat(filepath.Join(root, ".well-known", "geul")); err != nil {
		t.Skip("draft/.well-known/geul not present")
	}
	c, err := ConfigFromEnv(func(k string) string {
		return map[string]string{"GEULPROXY": "file://" + root + ",off"}[k]
	})
	if err != nil {
		t.Fatal(err)
	}
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	id, _ := core.EncodeDocID([]core.Word{0}, []core.Word{0xffff, 0, 0, 7})
	in := enc(id)
	for i, imp := range [][2]core.Word{{0x1018, 0}, {0x101C, 1}} {
		ptr, _ := core.EncodeDocPointer([]core.Word{0}, codecDocID(imp[0]), []core.Word{0, imp[1]})
		in = append(in, enc(core.Record{Kind: core.KindMeta, Codec: core.MetaImportDoc, Payload: append(enc(ptr), core.Word(i+1))})...)
	}
	in = append(append(in, 0x180E), enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("root")))...)
	g, err := (&Resolver{Config: c}).Resolve(context.Background(), core.PutWords(core.WrapDocument(append(in, 0x180F))))
	if err != nil {
		t.Fatal(err)
	}
	pub := g.Docs[NodeKey{DocKey: codecKey(0x1018), Version: 1}]
	field := g.Docs[NodeKey{DocKey: codecKey(0x101C), Version: 1}]
	if len(g.Docs) != 2 || pub == nil || field == nil || pub.Source != "file://"+root {
		t.Fatalf("got %+v", g.Docs)
	}
}
