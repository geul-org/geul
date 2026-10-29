//ff:func feature=pack type=helper control=sequence
//ff:what small — build options with tiny blocks (two entries per block) so the fixture gets internal levels; fixed clock, verify on
package pack

import "time"

// small forces P depth 2 and H depth 3 on the fixture.
func small() BuildOptions {
	return BuildOptions{PLeaf: 2, HLeaf: 2, HFanout: 2, HRoot: 2, Verify: true,
		Now: func() time.Time { return time.Unix(1_800_000_000, 0) }}
}
