//ff:func feature=toulmin type=encoder control=sequence
//ff:what TestDoCodec — DO without once, DO-ONCE with once
package toulmin

import "testing"

func TestDoCodec(t *testing.T) {
	if doCodec(false).Codec != CodecDo || doCodec(true).Codec != CodecDoOnce {
		t.Fatal("do codec mismatch")
	}
}
