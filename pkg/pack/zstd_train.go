//ff:func feature=pack type=encoder control=sequence
//ff:what Zstd.Train — a formatted zstd dictionary with id: content chosen by coverHistory, entropy tables fitted at z's level
package pack

import "github.com/klauspost/compress/zstd"

// Train builds a dictionary from samples. The content is chosen
// deterministically (coverHistory), so the same samples give the same bytes;
// klauspost's own trainer (package dict) does not. CompatV155 keeps it usable
// by zstd 1.5.5 and older.
func (z *Zstd) Train(samples [][]byte, id uint32) (*Dict, error) {
	b, err := zstd.BuildDict(zstd.BuildDictOptions{
		ID:         id,
		Contents:   samples[:min(len(samples), dictStatCount)],
		History:    coverHistory(samples, z.dictSize()),
		Offsets:    [3]int{1, 4, 8},
		CompatV155: true,
		Level:      z.level(),
	})
	if err != nil {
		return nil, err
	}
	return &Dict{ID: id, Bytes: b}, nil
}
