//ff:func feature=build type=renderer control=sequence
//ff:what IndexJSON — the listing of a tree directory as JSON for the tree browser, the same JSON serve answers
package server

import (
	"encoding/json"
	"io"
)

// IndexJSON writes the listing of the tree directory at URL path urlPath
// (ending in /) under root as the tree browser's JSON, for a static copy of
// the tree (build).
func IndexJSON(w io.Writer, root, urlPath string) error {
	p, err := treeIndex(root, urlPath)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(p)
}
