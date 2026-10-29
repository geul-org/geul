//ff:type feature=pack type=model
//ff:what fileSys — the writer's file operations (temp file, rename, remove, mkdir), replaceable in tests
package pack

// fileSys is the output side of a writer.
type fileSys interface {
	MkdirAll(dir string) error
	CreateTemp(dir string) (objFile, error)
	Rename(from, to string) error
	Remove(name string) error
}
