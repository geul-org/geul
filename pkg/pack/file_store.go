//ff:type feature=pack type=resolver
//ff:what FileStore — a Store over a local directory holding latest.txt, m/ and o/ (pread, standard library only)
package pack

// FileStore reads objects under Root with os.ReadFile and File.ReadAt.
type FileStore struct {
	Root string
}
