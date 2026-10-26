//ff:func feature=cli type=helper control=sequence
//ff:what harness.doc — test fixture: TRIPLE's spec, a local Issuer "me" and document me/0x…01 published twice (v2 changes #2 and adds #3)
package cli

import "path/filepath"

// doc publishes version 1 (LABEL #1, TRIPLE #2) and version 2 (#2 changed,
// ALIAS #3 added) of me/0x0000000000000001 and returns its address and the
// draft path.
func (h *harness) doc() (string, string) {
	addr := "me/0x0000000000000001"
	draft := filepath.Join(h.data, "d.geul")
	h.tripleSpec()
	h.must("issuer", "new", "--name", "me")
	h.must("new", addr, "-o", draft)
	h.must("add", draft, "LABEL", "Q42", `"Douglas Adams"@en`)
	h.must("add", draft, "TRIPLE", "P31", "Q42", "Q5", "normal")
	h.must("publish", draft)
	h.must("set", draft, "#2", "P31", "Q42", "Q36180", "normal")
	h.must("add", draft, "ALIAS", "Q42", `"DNA"@en`)
	h.must("publish", draft)
	return addr, draft
}
