//ff:func feature=server type=handler control=sequence
//ff:what serveDoc — a document path without .geul: its bytes when Accept asks for GEUL, the viewer page otherwise
package server

import (
	"net/http"
	"path/filepath"
)

// serveDoc answers the extensionless path p of the document p.geul (GEUL-ORG
// 형식 지정): a tool that accepts application/vnd.geul gets the bytes, a
// browser gets the viewer page with the tree browser on. The answer depends
// on Accept, so it says so in Vary.
func serveDoc(w http.ResponseWriter, r *http.Request, root, p string) {
	file := filepath.Join(root, filepath.FromSlash(p)) + ".geul"
	w.Header().Set("Vary", "Accept")
	if acceptsGeul(r) {
		serveRaw(w, r, file)
		return
	}
	renderFile(w, file, titleOf(root, p), true, treeSpec(root))
}
