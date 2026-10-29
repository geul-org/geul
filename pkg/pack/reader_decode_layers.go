//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what Reader.decodeLayers — decode each layer root in place, checking it belongs to the right key space
package pack

import "context"

// decodeLayers unframes and decodes every root of ls; p says the P space.
func (r *Reader) decodeLayers(ctx context.Context, ls []layer, p bool) error {
	for i := range ls {
		raw, err := r.unstore(ctx, ls[i].enc, ls[i].root, true)
		if err != nil {
			return err
		}
		if ls[i].blk, err = decodeBlock(raw); err != nil {
			return err
		}
		if isPKind(ls[i].blk.kind) != p {
			return corrupt("layer root in the wrong key space")
		}
	}
	return nil
}
