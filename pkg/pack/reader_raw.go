//ff:func feature=pack type=resolver control=sequence
//ff:what Reader.raw — the stored bytes at a Loc after checking object number, object kind and bounds
package pack

import "context"

// raw reads l's stored bytes with one range read. The object must exist in
// OBJECTS with the expected kind and l must lie after its header.
func (r *Reader) raw(ctx context.Context, kind byte, l Loc) ([]byte, error) {
	if uint64(l.Obj) >= uint64(len(r.m.objects)) {
		return nil, corrupt("object number out of range")
	}
	o := r.m.objects[l.Obj]
	if o.kind != kind || l.Off < headerSize || l.Off > o.size || l.Len > o.size-l.Off {
		return nil, corrupt("location outside its object")
	}
	return r.readRange(ctx, o, l.Off, l.Len)
}
