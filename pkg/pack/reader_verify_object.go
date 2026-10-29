//ff:func feature=pack type=rule control=sequence
//ff:what Reader.verifyObject — an object's header against OBJECTS, its SHA-256 against its name, no bytes past its size
package pack

import (
	"context"
	"fmt"
)

// verifyObject checks object i end to end.
func (r *Reader) verifyObject(ctx context.Context, i int) error {
	o := r.m.objects[i]
	hb, err := r.readRange(ctx, o, 0, headerSize)
	if err != nil {
		return err
	}
	h, err := parseObjectHeader(hb)
	if err == nil && (h.kind != o.kind || h.size != o.size) {
		err = corrupt(objectName(o) + " header disagrees with OBJECTS")
	}
	if err != nil {
		return err
	}
	sum, err := r.objectSum(ctx, o)
	if err == nil && sum != o.hash {
		err = corrupt(objectName(o) + " SHA-256 is not its name")
	}
	if err != nil {
		return err
	}
	if _, err := r.s.ReadRange(ctx, objectName(o), o.size, 1); err == nil {
		return corrupt(fmt.Sprintf("%s is longer than %d bytes", objectName(o), o.size))
	}
	return nil
}
