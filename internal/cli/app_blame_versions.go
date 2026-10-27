//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what app.blameVersions — the live versions of a document up to the asked one (all for none), oldest first, as snapshots
package cli

import (
	"errors"

	"github.com/geul-org/gogeul/pkg/delta"
)

// blameVersions reads each version that is not deleted, up to addr's version
// (all when 0).
func (a *app) blameVersions(addr string) ([]delta.Version, error) {
	k, err := a.key(addr)
	if err != nil {
		return nil, err
	}
	d, err := a.dir()
	if err != nil {
		return nil, err
	}
	h, err := d.Log(k.Issuer, k.Doc)
	if err != nil {
		return nil, err
	}
	var vs []delta.Version
	for _, e := range h.Entries {
		if e.Deleted || (k.Version != 0 && e.Version > k.Version) {
			continue
		}
		b, err := d.At(k.Issuer, k.Doc, e.Version)
		if err != nil {
			return nil, err
		}
		vs = append(vs, delta.Version{V: e.Version, Doc: b})
	}
	if len(vs) == 0 {
		return nil, errors.New("blame: no live version")
	}
	return vs, nil
}
