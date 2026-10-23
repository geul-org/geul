//ff:type feature=resolve type=resolver
//ff:what Memory — a Source backed by a path → bytes map (tests)
package resolve

// Memory is a Source over an in-memory map from path to bytes.
type Memory map[string][]byte
