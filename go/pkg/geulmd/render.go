//ff:func feature=geulmd type=encoder control=sequence
//ff:what Render — a .geul document as GEULMD: a `geul` HEAD block then a `geul` BODY block
package geulmd

import (
	"github.com/geul-org/geul/go/pkg/core"
)

// Render converts b, exactly one GEUL document, to GEULMD text: the first fenced
// block (info string `geul`) holds the HEAD — and marks the file as GEULMD, not
// natural language — and the second holds the BODY, one record per line. The
// DOC, BODY and END structure records are not written; they are recomputed on
// encoding.
func Render(b []byte) (string, error) {
	ws, err := core.ReadWords(b)
	if err != nil {
		return "", err
	}
	d, err := core.SplitDocument(ws)
	if err != nil {
		return "", err
	}
	head, ctx := renderHead(d.Head)
	body := renderBody(d.Body, ctx)
	return block(head) + "\n" + block(body), nil
}
