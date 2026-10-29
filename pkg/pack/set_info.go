//ff:type feature=pack type=model
//ff:what SetInfo — one pack set of an Archive as reported to users: name, directory, generation, manifest, objects
package pack

// SetInfo describes a set ("" Name is the set at the top).
type SetInfo struct {
	Name       string
	Dir        string
	Generation uint64
	Manifest   string
	Objects    int
}
