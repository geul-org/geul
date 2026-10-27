//ff:func feature=cli type=command control=sequence
//ff:what app.newDraft — write a new document's draft; refused when the document is already published
package cli

import (
	"errors"
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// newDraft resolves addr (GEULISSUER when it has no Issuer part), refuses a
// document the data folder already publishes (open it instead), and writes
// HEAD SPEC-VERSION, DOC-ID and, for a local key's Issuer, ISSUER-KEY, with
// an empty BODY. key has read the canonical Issuer and document ID and the
// local keys, so encoding DOC-ID and reading the keys again do not fail.
func (a *app) newDraft(addr, out string, force bool) error {
	if !strings.Contains(addr, "/") {
		addr = a.cfg.Get("GEULISSUER") + "/" + addr
	}
	k, err := a.key(addr)
	if err != nil {
		return err
	}
	if d, err := a.dir(); err == nil {
		if _, err := d.Log(k.Issuer, k.Doc); err == nil {
			return errors.New("the document is published; use geul open to edit its next version")
		}
	}
	id, _ := core.EncodeDocID(k.Issuer, k.Doc)
	keys, _ := a.localKeys()
	b, _ := core.EncodeDocument(issuerKeyHead(keys, k.Issuer, []core.Record{core.EncodeSpecVersion(core.SpecVersion), id}), nil)
	return writeNew(out, b, force)
}
