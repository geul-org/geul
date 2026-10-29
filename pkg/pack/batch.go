//ff:type feature=pack type=model
//ff:what Batch — the change a new generation records: paths to put (bytes from Get), paths to delete, or a full listing
package pack

// Batch is one generation's input. Put paths whose bytes equal the previous
// generation's are left as they are. Full means Put lists the whole archive,
// so every previous path not in Put is deleted. Get returns a Put path's bytes
// and is called again while writing; the bytes must not change in between.
type Batch struct {
	Put    []string
	Delete []string
	Full   bool
	Get    func(path string) ([]byte, error)
}
