//ff:func feature=pack type=encoder control=sequence
//ff:what Zstd.encoder — the encoder for no dictionary or for d, made on first use and kept
package pack

import "github.com/klauspost/compress/zstd"

// encoder returns z's plain encoder (d nil) or d's.
func (z *Zstd) encoder(d *Dict) (*zstd.Encoder, error) {
	z.mu.Lock()
	defer z.mu.Unlock()
	slot := &z.plain
	if d != nil {
		slot = &d.enc
	}
	if *slot == nil {
		e, err := newEncoder(z.level(), z.workers(), d)
		if err != nil {
			return nil, err
		}
		*slot = e
	}
	return *slot, nil
}
