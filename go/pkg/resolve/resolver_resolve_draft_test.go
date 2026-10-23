//ff:func feature=resolve type=resolver control=sequence
//ff:what TestResolverResolveDraft — the draft/t002/.well-known/geul samples resolve through a custom GEULPROXY file:// source
package resolve

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

// TestResolverResolveDraft reads the sample documents published in the GEUL spec
// repository at draft/t002/.well-known/geul/ (the temporary stand-in for
// geul.org/.well-known/geul/). It is skipped when this module is checked out on
// its own, without the spec repository around it.
func TestResolverResolveDraft(t *testing.T) {
	root, _ := filepath.Abs("../../../..")
	if _, err := os.Stat(filepath.Join(root, ".well-known", "geul")); err != nil {
		t.Skip("draft/t002/.well-known/geul not present")
	}
	c, err := ConfigFromEnv(func(k string) string {
		return map[string]string{"GEULPROXY": "file://" + root + ",off"}[k]
	})
	if err != nil {
		t.Fatal(err)
	}
	g, err := (&Resolver{Config: c}).Resolve(context.Background(), sampleDoc(7, "root", [3]uint32{1, 3, 0}))
	if err != nil {
		t.Fatal(err)
	}
	a := g.Docs[sampleKey(1, 0).DocKey()]
	if len(g.Docs) != 3 || a.Version != 2 || a.Source != "file://"+root {
		t.Fatalf("got %+v", g.Docs)
	}
}
