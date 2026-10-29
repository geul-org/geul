//ff:func feature=pack type=resolver control=selection
//ff:what Reader.applyEntry — apply one inventory entry to a state: set (and h/ share) or delete (§7.1)
package pack

import (
	"context"
	"fmt"
)

// applyEntry applies e of the data pack obj to st.
func (r *Reader) applyEntry(ctx context.Context, st map[string]Loc, obj uint32, inv inventory, e invEntry) error {
	l := Loc{Obj: obj, Off: e.off, Len: uint64(e.n), Enc: e.enc}
	switch e.kind {
	case invP, invNamed:
		p, err := invPath(inv, e)
		if err != nil {
			return err
		}
		st[p] = l
		return r.applyShare(ctx, st, e, l)
	case invH:
		p, err := r.hashPath(ctx, l)
		st[p] = l
		return err
	case invPDel, invNamedDel:
		p, err := invPath(inv, e)
		delete(st, p)
		return err
	case invDict:
		return nil
	default:
		return corrupt(fmt.Sprintf("inventory entry kind %d", e.kind))
	}
}
