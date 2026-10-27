//ff:func feature=server type=helper control=sequence
//ff:what treeIndex — the listing data of the tree directory at a URL path, its entries named by treeLabels
package server

import (
	"path/filepath"
	"strings"
)

// treeIndex lists the tree directory at URL path urlPath (ending in /) under
// root, showing Issuer and document IDs by their names.
func treeIndex(root, urlPath string) (indexPage, error) {
	dir := filepath.Join(root, filepath.FromSlash(urlPath))
	return indexPageOf(dir, urlPath, treeLabels(root, strings.TrimSuffix(urlPath, "/")))
}
