//ff:func feature=cli type=command control=iteration dimension=1
//ff:what app.resolve — resolve a document's IMPORT graph and print one line per bound version and its source
package cli

import (
	"context"
	"fmt"
	"sort"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// resolve builds the dependency graph of the document (resolve.Resolver)
// with the data folder as the first source, then GEULPACK, GEULPROXY and
// GEULPRIVATE.
// Each imported version is printed as `<Issuer>/<문서>@<버전> <출처>`.
func (a *app) resolve(arg string) error {
	b, err := a.readDoc(arg)
	if err != nil {
		return err
	}
	cfg, err := resolve.ConfigFromEnv(a.cfg.Get)
	if err == nil {
		cfg, err = packSources(cfg, a.cfg.Get("GEULPACK"))
	}
	if err != nil {
		return err
	}
	if d, err := a.dir(); err == nil {
		cfg.Proxy = append([]resolve.Entry{{Name: "file://" + d.Root, Source: d}}, cfg.Proxy...)
	}
	r := resolve.Resolver{Config: cfg}
	g, err := r.Resolve(context.Background(), b)
	if err != nil {
		return err
	}
	lines := []string{}
	for k, n := range g.Docs {
		if n != g.Root {
			lines = append(lines, fmt.Sprintf("%s/%s@%d %s", k.Issuer, k.Doc, n.Version, n.Source))
		}
	}
	sort.Strings(lines)
	if done, err := a.json(lines); done {
		return err
	}
	for _, l := range lines {
		fmt.Fprintln(a.out, l)
	}
	return nil
}
