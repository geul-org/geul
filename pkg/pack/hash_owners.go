//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what hashOwners — content hash → Issuer for every Issuer path that is a GEUL document, so h/ copies follow their documents
package pack

import "os"

// hashOwners reads the Issuer paths of the tree at base (resolve.Dir layout).
func hashOwners(base string, paths []string) (map[string]string, error) {
	out := map[string]string{}
	for _, p := range paths {
		owner := ownerOf(p)
		if owner == "" {
			continue
		}
		b, err := os.ReadFile(treeFile(base, p))
		if err != nil {
			return nil, err
		}
		if h, ok := contentHash(b); ok {
			out[string(h[:])] = owner
		}
	}
	return out, nil
}
