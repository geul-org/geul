//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what app.localKeys — every local Issuer key in name order; none when the keys folder does not exist
package cli

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// localKeys reads every `*.pem` of keysDir (readKey).
func (a *app) localKeys() ([]localKey, error) {
	ents, err := os.ReadDir(a.keysDir())
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var out []localKey
	for _, e := range ents {
		if !strings.HasSuffix(e.Name(), ".pem") {
			continue
		}
		k, err := readKey(filepath.Join(a.keysDir(), e.Name()))
		if err != nil {
			return nil, err
		}
		out = append(out, k)
	}
	return out, nil
}
