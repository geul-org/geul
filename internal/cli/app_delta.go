//ff:func feature=cli type=command control=sequence
//ff:what app.delta — read two versions of a document and write their diff document
package cli

import (
	"errors"
	"os"
	"strings"

	"github.com/geul-org/gogeul/pkg/delta"
)

// delta reads base (`<문서>@v1`) and `@v2` of the same document and writes
// delta.Diff of them to out. Diff refuses a pair whose order a diff cannot
// keep (new Edges must come after the base's).
func (a *app) delta(base, next, out string) error {
	addr, _, _ := strings.Cut(base, "@")
	if !strings.HasPrefix(next, "@") {
		return errors.New("delta: the second argument is @VERSION of the same document")
	}
	k, err := a.key(base)
	if err != nil {
		return err
	}
	b, n, err := a.readPair(base, addr+next)
	if err != nil {
		return err
	}
	d, err := delta.Diff(b, n, k.Version)
	if err != nil {
		return err
	}
	return os.WriteFile(out, d, 0o644)
}
