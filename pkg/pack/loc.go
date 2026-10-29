//ff:type feature=pack type=model
//ff:what Loc — where a record or directory block lies: object number, offset, stored length, encoding
package pack

// Loc locates stored bytes: Obj is the position in the manifest's OBJECTS,
// Off the absolute offset in that object, Len the stored length and Enc the
// encoding (0 raw, 1 zstd, 0xFF deletion marker).
type Loc struct {
	Obj uint32
	Off uint64
	Len uint64
	Enc byte
}
