//ff:type feature=pack type=model
//ff:what layer — one directory tree of P or H: entry count, root block as stored and decoded
package pack

// layer is a P or H directory tree whose root lives in the manifest.
type layer struct {
	count uint64
	enc   byte
	root  []byte // as stored (encoding enc)
	blk   block  // decoded root
}
