//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what findSets — the pack sets in a `.well-known/geul/` directory: the top one, then Issuer folders with latest.txt, in Dir order
package pack

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
)

// findSets lists the sets under wk; a missing wk has none.
func findSets(wk string) ([]*packSet, error) {
	var out []*packSet
	if _, err := os.Stat(filepath.Join(wk, latestFile)); err == nil {
		out = append(out, &packSet{dir: wk})
	}
	ents, err := os.ReadDir(wk)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return nil, err
	}
	var ids []string
	for _, e := range ents {
		if _, err := os.Stat(filepath.Join(wk, e.Name(), latestFile)); err == nil && isIssuer(e.Name()) {
			ids = append(ids, e.Name())
		}
	}
	sort.Slice(ids, func(i, j int) bool { return issuerLess(ids[i], ids[j]) })
	for _, id := range ids {
		out = append(out, &packSet{name: id, dir: filepath.Join(wk, id)})
	}
	return out, nil
}
