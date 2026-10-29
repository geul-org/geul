//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what keepData — the data packs of OBJECTS in order, with old → new numbers (index objects dropped)
package pack

// keepData keeps the data packs (their order is the layer order, §4.3) and
// maps their old numbers to the new ones.
func keepData(objs []object) ([]object, map[uint32]uint32) {
	var out []object
	remap := map[uint32]uint32{}
	for i, o := range objs {
		if o.kind == kindData {
			remap[uint32(i)] = uint32(len(out))
			out = append(out, o)
		}
	}
	return out, remap
}
