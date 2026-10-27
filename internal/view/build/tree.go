//ff:func feature=build type=command control=sequence
//ff:what Tree — bake a .well-known/geul tree into static files: bytes, viewer pages, listings and scripts
package build

import (
	"path/filepath"

	"github.com/geul-org/gogeul/internal/view/static"
	"github.com/geul-org/gogeul/pkg/resolve"
)

// Tree writes a static copy of the tree under root/.well-known/geul to
// dst/.well-known/geul, at the paths of the public address space (GEUL-ORG
// 형식 지정), for a plain web host:
//
//	<path>.geul       the document's bytes (application/vnd.geul)
//	<path>            its viewer page (text/html; the host picks by Accept)
//	<dir>/index.html  the directory listing
//	<dir>/index.json  the listing for the tree browser
//	_static/          the graph scripts
//
// The pages are the ones serve answers. The codec specs come from the tree;
// a tree whose specs do not read lays records out generically, as serve does.
// A document that does not decode stops the build.
func Tree(root, dst string) (Count, error) {
	b := &baker{root: root, src: filepath.Join(root, ".well-known", "geul"), out: filepath.Join(dst, ".well-known", "geul")}
	if err := checkOut(b.src, b.out); err != nil {
		return Count{}, err
	}
	b.spec, _ = resolve.Dir{Root: root}.CodecSpec()
	if err := filepath.WalkDir(b.src, b.visit); err != nil {
		return b.count, err
	}
	return b.count, static.Write(filepath.Join(b.out, "_static"))
}
