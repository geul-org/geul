//ff:func feature=pack type=writer control=sequence
//ff:what TestKeepData — data packs keep their order, index objects drop, numbers remap
package pack

import "testing"

func TestKeepData(t *testing.T) {
	objs := []object{{kind: kindData, size: 1}, {kind: kindIndex}, {kind: kindData, size: 2}}
	out, remap := keepData(objs)
	if len(out) != 2 || out[1].size != 2 || remap[2] != 1 || remap[0] != 0 {
		t.Fatal(out, remap)
	}
}
