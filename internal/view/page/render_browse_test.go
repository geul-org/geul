//ff:func feature=page type=renderer control=iteration dimension=1
//ff:what TestRenderBrowse — the 탐색 menu, its sidebar and script appear only when Browse is set
package page

import (
	"bytes"
	"strings"
	"testing"
)

func TestRenderBrowse(t *testing.T) {
	var on, off bytes.Buffer
	Render(&on, Page{Title: "t", Browse: true})
	Render(&off, Page{Title: "t"})
	for _, want := range []string{`id="browseBtn"`, `<aside id="browse">`, "fetch(u+'index.json')"} {
		if !strings.Contains(on.String(), want) || strings.Contains(off.String(), want) {
			t.Fatalf("%q: browse on/off wrong", want)
		}
	}
}
