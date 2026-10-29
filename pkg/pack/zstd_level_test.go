//ff:func feature=pack type=helper control=sequence
//ff:what TestZstdLevel — 0 is level 19 (best); low levels map to the fastest speed
package pack

import (
	"testing"

	"github.com/klauspost/compress/zstd"
)

func TestZstdLevel(t *testing.T) {
	if (&Zstd{}).level() != zstd.SpeedBestCompression || (&Zstd{Level: 1}).level() != zstd.SpeedFastest || (&Zstd{Level: 3}).level() != zstd.SpeedDefault {
		t.Fatal("level")
	}
}
