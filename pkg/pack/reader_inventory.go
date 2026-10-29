//ff:func feature=pack type=resolver control=sequence
//ff:what Reader.inventory — a data pack's header and checksummed inventory, read by range
package pack

import "context"

// inventory reads object obj's header and inventory (§4.4, §4.5).
func (r *Reader) inventory(ctx context.Context, obj uint32) (inventory, error) {
	o := r.m.objects[obj]
	hb, err := r.readRange(ctx, o, 0, headerSize)
	if err != nil {
		return inventory{}, err
	}
	h, err := parseObjectHeader(hb)
	if err == nil && (h.kind != kindData || h.size != o.size) {
		err = corrupt("data pack header disagrees with OBJECTS")
	}
	var raw []byte
	if err == nil {
		raw, err = r.readRange(ctx, o, h.invOff, h.invLen)
	}
	if err == nil {
		raw, err = r.unzstd(ctx, raw, true)
	}
	if err != nil {
		return inventory{}, err
	}
	return parseInventory(raw)
}
