//ff:type feature=pack type=model
//ff:what block — a decoded directory block: kind (P/H leaf/internal), ascending keys and their locations
package pack

// block is a directory block (§4.7). keys are 16-byte P keys or 32-byte H
// keys in strictly ascending byte order; locs[i] is the record (leaf) or child
// block (internal) of keys[i].
type block struct {
	kind byte
	keys []string
	locs []Loc
}
