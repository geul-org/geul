//ff:type feature=pack type=model
//ff:what fsFuncs — a fileSys whose operations fail on demand, otherwise the real file system
package pack

// fsFuncs fails the named operation after the given number of successful calls
// (-1: never). create may wrap the file it returns.
type fsFuncs struct {
	mkdir, create, rename int
	wrap                  func(objFile) objFile
}
