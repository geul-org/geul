//ff:type feature=pack type=model
//ff:what manifest — one generation: chain fields, OBJECTS, ISSUERS, DICTS, NAMED and the P and H layers
package pack

// manifest is a parsed `.gpm` (§4.3). Layers are newest first.
type manifest struct {
	gen     uint64
	created int64
	prev    [32]byte
	objects []object
	issuers []string // canonical hex; the number is the position
	dicts   []dictEntry
	named   []namedEntry // ascending by path
	pLayers []layer
	hLayers []layer
}
