//ff:func feature=pack type=helper control=sequence
//ff:what TestBuildOptionsWithDefaults — zero values become the recommended parameters; set values stay
package pack

import (
	"testing"
	"time"
)

func TestBuildOptionsWithDefaults(t *testing.T) {
	o := BuildOptions{}.withDefaults()
	if o.PackSize != DefaultPackSize || o.PLeaf != DefaultPLeaf || o.HLeaf != DefaultHLeaf || o.HFanout != DefaultHFanout || o.HRoot != DefaultHRoot || o.HLayers != DefaultHLayers || o.Codec == nil || o.Now == nil {
		t.Fatal(o)
	}
	now := func() time.Time { return time.Time{} }
	o = BuildOptions{PackSize: 5, HFanout: 1, Codec: Stored{}, Now: now}.withDefaults()
	if o.PackSize != 5 || o.HFanout != DefaultHFanout || o.Now().Year() != 1 {
		t.Fatal(o)
	}
}
