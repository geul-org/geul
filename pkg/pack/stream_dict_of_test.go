//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestStreamDictOf — enough similar samples give a dictionary with the first ID; Stored, DictSize < 0, too few samples, failed training and samples with no saving give none
package pack

import (
	"crypto/sha256"
	"testing"
)

func TestStreamDictOf(t *testing.T) {
	o := streamTestOptions().withDefaults()
	if d := streamDictOf(o); d == nil || d.ID != FirstDictID {
		t.Fatalf("%v", d)
	}
	for name, c := range map[string]Codec{"stored": Stored{}, "no dict": &Zstd{DictSize: -1}} {
		o2 := o
		o2.Build.Codec = c
		if streamDictOf(o2) != nil {
			t.Errorf("%s: want none", name)
		}
	}
	few := o
	few.Samples = o.Samples[:dictMinSamples-1]
	if streamDictOf(few) != nil {
		t.Error("few samples: want none")
	}
	tiny := o
	tiny.Samples = make([][]byte, 40) // all empty: nothing to train on
	if streamDictOf(tiny) != nil {
		t.Error("empty samples: want none")
	}
	random := o
	random.Samples = nil
	for i := 0; i < 40; i++ {
		var b []byte
		for j := 0; j < 64; j++ {
			h := sha256.Sum256([]byte{byte(i), byte(j)})
			b = append(b, h[:]...)
		}
		random.Samples = append(random.Samples, b)
	}
	if streamDictOf(random) != nil {
		t.Error("random samples: want none")
	}
}
