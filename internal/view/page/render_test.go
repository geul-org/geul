//ff:func feature=page type=renderer control=iteration dimension=1
//ff:what TestRender — the page HTML carries rows, labels, links, spans, the SOURCE and graph as JSON and the graph scripts; a failing writer errors
package page

import (
	"bytes"
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	var b bytes.Buffer
	if err := Render(&b, Build(fixtureDoc(t), "x.geul")); err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"<title>x.geul · geul</title>", `class="item node" id="e1"`, `data-off="2" data-len="6"`,
		"<b>사안</b>", `<a href="#e1">#1 사안</a>`, `"# 사안\n규칙"`, `class="item raw" id="e16"`,
		`"source":"e2","target":"e1","kind":"ATTACKS"`, `<script src="/.well-known/geul/_static/cytoscape.min.js">`} {
		if !strings.Contains(b.String(), want) {
			t.Fatalf("missing %q", want)
		}
	}
	if err := Render(failWriter{}, Page{}); err == nil {
		t.Fatal("failing writer: want error")
	}
}
