//ff:func feature=build type=renderer control=sequence
//ff:what TestIndexHTML — the listing page links documents without .geul; a missing directory fails
package server

import (
	"bytes"
	"strings"
	"testing"
)

func TestIndexHTML(t *testing.T) {
	root := treeFixture(t)
	var b bytes.Buffer
	if err := IndexHTML(&b, treeOf(root), "/.well-known/geul/h/"); err != nil || !strings.Contains(b.String(), `href="/.well-known/geul/h/a"`) || !strings.Contains(b.String(), `href="/.well-known/geul/"`) {
		t.Fatalf("%v %s", err, b.String())
	}
	if err := IndexHTML(&b, treeOf(root), "/.well-known/geul/none/"); err == nil {
		t.Fatal("missing: want error")
	}
}
