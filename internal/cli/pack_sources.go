//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what packSources — resolver sources with pack archives: file:// folders holding pack sets read through them, GEULPACK sets first
package cli

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/geul-org/gogeul/pkg/pack"
	"github.com/geul-org/gogeul/pkg/resolve"
)

// packSources rewrites cfg.Proxy: a `file://` folder whose `.well-known/geul/`
// holds pack sets is read through them (the notation is unchanged), and the
// comma-separated GEULPACK items come before the GEULPROXY ones.
func packSources(cfg resolve.Config, packs string) (resolve.Config, error) {
	for i, e := range cfg.Proxy {
		d, ok := e.Source.(resolve.Dir)
		wk := filepath.Join(d.Root, ".well-known", "geul")
		if !ok || d.Store != nil || !pack.IsPacked(wk) {
			continue
		}
		a, err := pack.OpenDir(context.Background(), wk, packOptions())
		if err != nil {
			return cfg, err
		}
		cfg.Proxy[i].Source = resolve.Dir{Root: d.Root, Store: a}
	}
	var pre []resolve.Entry
	for _, item := range strings.Split(packs, ",") {
		if item = strings.TrimSpace(item); item == "" {
			continue
		}
		e, err := packEntry(item)
		if err != nil {
			return cfg, err
		}
		pre = append(pre, e)
	}
	cfg.Proxy = append(pre, cfg.Proxy...)
	return cfg, nil
}
