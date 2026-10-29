//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what memGen — a memStore holding one manifest (and latest.txt) for crafted manifests; returns the store and manifest name
package pack

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// memGen stores m as generation m.gen with the extra objects given.
func memGen(m *manifest, objs map[string][]byte) (memStore, string) {
	b := encodeManifest(m)
	sum := sha256.Sum256(b)
	name := hex.EncodeToString(sum[:])
	s := memStore{"m/" + name + ".gpm": b, latestFile: []byte(fmt.Sprintf("%d %s\n", m.gen, name))}
	for k, v := range objs {
		s[k] = v
	}
	return s, name
}
