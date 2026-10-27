//ff:func feature=build type=renderer control=sequence
//ff:what IndexHTML — the listing page of a tree directory, the same page serve answers
package server

import "io"

// IndexHTML writes the listing page of the tree directory at URL path urlPath
// (ending in /) under root, for a static copy of the tree (build).
func IndexHTML(w io.Writer, root, urlPath string) error {
	p, err := treeIndex(root, urlPath)
	if err != nil {
		return err
	}
	return indexTemplate.Execute(w, p)
}
