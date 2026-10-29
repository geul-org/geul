//ff:type feature=pack type=model
//ff:what put — a path to write: its key space, SHA-256, size and content hash, the h/ path it also serves, its location
package pack

// put is one changed path. A shared h/ put has via set to the record that
// serves it; other puts get their own record at loc.
type put struct {
	path  string
	ref   pathRef
	sum   [32]byte
	n     int
	ch    [32]byte
	chOK  bool
	share string
	via   *put
	loc   Loc
}
