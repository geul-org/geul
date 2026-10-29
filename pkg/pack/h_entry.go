//ff:type feature=pack type=model
//ff:what hEntry — one H key (content hash) and the record location serving it, as spilled in a run
package pack

// hEntry is an H key and the location of the record that serves it. A run
// file holds them sorted by key, then location, hEntrySize bytes each.
type hEntry struct {
	key [32]byte
	loc Loc
}
