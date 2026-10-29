//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what splitPaths — group an archive's paths into Issuer sets: Issuer paths by Issuer, h/ copies with the Issuer of the document they copy, the rest home
package pack

// splitPaths assigns every path to a set name (Issuer hex, or "" when the
// archive has no Issuer paths).
func splitPaths(base string, paths []string) (map[string][]string, error) {
	hs, err := hashOwners(base, paths)
	if err != nil {
		return nil, err
	}
	seen := map[string]bool{}
	for _, owner := range hs {
		seen[owner] = true
	}
	for _, p := range paths {
		if o := ownerOf(p); o != "" {
			seen[o] = true
		}
	}
	home := splitHome(seen)
	out := map[string][]string{}
	for _, p := range paths {
		owner := ownerOf(p)
		if ref := parsePath(p); owner == "" && ref.space == spaceH {
			owner = hs[ref.hash]
		}
		if owner == "" {
			owner = home
		}
		out[owner] = append(out[owner], p)
	}
	return out, nil
}
