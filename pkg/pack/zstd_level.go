//ff:func feature=pack type=helper control=sequence
//ff:what Zstd.level — the klauspost speed for the zstd level (0 is DefaultZstdLevel)
package pack

import "github.com/klauspost/compress/zstd"

// level maps z.Level to an EncoderLevel.
func (z *Zstd) level() zstd.EncoderLevel {
	return zstd.EncoderLevelFromZstd(defaultInt(z.Level, 1, DefaultZstdLevel))
}
