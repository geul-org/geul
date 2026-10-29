//ff:func feature=server type=helper control=sequence
//ff:what treeOf — the directory Store of a test root (root/.well-known/geul)
package server

import "github.com/geul-org/gogeul/pkg/resolve"

// treeOf opens root's tree the way the CLI does for a directory archive.
func treeOf(root string) resolve.Store {
	return resolve.Dir{Root: root}.Tree()
}
