//ff:func feature=cli type=helper control=sequence
//ff:what app.readDoc — a document's bytes from a file path, else from the data folder by address (diff-stored versions rebuilt)
package cli

import "os"

// readDoc reads arg as a file when one exists there, else resolves it as
// `<Issuer>/<문서>[@<버전>]` and reads that version (resolve.Dir.At, 0 is
// the newest live version).
func (a *app) readDoc(arg string) ([]byte, error) {
	if fi, err := os.Stat(arg); err == nil && !fi.IsDir() {
		return os.ReadFile(arg)
	}
	k, err := a.key(arg)
	if err != nil {
		return nil, err
	}
	d, err := a.dir()
	if err != nil {
		return nil, err
	}
	return d.At(k.Issuer, k.Doc, k.Version)
}
