//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterReuseDict — a carried dictionary's bytes come from the previous generation; an unknown ID fails
package pack

import (
	"context"
	"testing"
)

func TestWriterReuseDict(t *testing.T) {
	out, _, _ := zstdSet(t, 100, &Zstd{DictSize: 4096})
	prev := opened(t, out)
	w := &writer{ctx: context.Background(), prev: prev}
	if err := w.reuseDict(FirstDictID); err != nil || w.dict.ID != FirstDictID || len(w.dict.Bytes) == 0 {
		t.Fatal(err)
	}
	if err := w.reuseDict(7); err == nil {
		t.Fatal("unknown")
	}
}
