//ff:func feature=cli type=helper control=sequence
//ff:what harness.packed — test fixture: the harness's directory archive converted to pack sets in a new data folder the harness then uses
package cli

import (
	"path/filepath"
)

// packed runs `geul pack build` into a fresh folder and moves the harness
// there with GEULSTORE unset, so later commands read and write pack sets.
func (h *harness) packed() {
	h.t.Helper()
	out := filepath.Join(h.t.TempDir(), "data")
	h.must("pack", "build", h.data, "-o", out)
	h.data, h.cwd = out, out
	delete(h.env, "GEULSTORE")
}
