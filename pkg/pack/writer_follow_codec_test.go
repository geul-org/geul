//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterFollowCodec — new and raw sets stay Stored; a dictionary or a compressed record means Zstd (also on Fresh); a broken index fails on Fresh
package pack

import (
	"context"
	"os"
	"testing"
)

func TestWriterFollowCodec(t *testing.T) {
	ctx := context.Background()
	follow := func(out string, fresh bool) (Codec, error) {
		o := small()
		o.Fresh = fresh
		w := &writer{ctx: ctx, out: out, o: o.withDefaults()}
		if err := w.open(); err != nil {
			return nil, err
		}
		return w.followCodec()
	}
	if c, err := follow(t.TempDir(), false); err != nil || c != (Stored{}) {
		t.Fatal(c, err)
	}
	raw, _, _ := zstdSet(t, 40, Stored{})
	if c, err := follow(raw, false); err != nil || c != (Stored{}) {
		t.Fatal(c, err)
	}
	withDict, _, _ := zstdSet(t, 100, &Zstd{DictSize: 4096})
	if c, _ := follow(withDict, false); c.(*Zstd) == nil {
		t.Fatal(c)
	}
	noDict, _, _ := zstdSet(t, 40, &Zstd{DictSize: -1})
	if c, _ := follow(noDict, false); c.(*Zstd) == nil {
		t.Fatal(c)
	}
	if c, err := follow(noDict, true); err != nil || c.(*Zstd) == nil {
		t.Fatal(c, err)
	}
	_, fx := built(t, small())
	os.Remove(objPath(t, opened(t, fx), fx, kindIndex))
	if _, err := follow(fx, true); err == nil {
		t.Fatal("entries")
	}
}
