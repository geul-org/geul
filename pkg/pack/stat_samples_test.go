//ff:func feature=pack type=encoder control=iteration dimension=1
//ff:what TestStatSamples — at most dictStatCount samples, each cut to one block; training on a sample bigger than a block no longer panics
package pack

import (
	"bytes"
	"testing"
)

func TestStatSamples(t *testing.T) {
	big := bytes.Repeat([]byte("GEUL codec spec ENUM BIND "), 10000) // 260 KB
	var ss [][]byte
	for i := 0; i < dictStatCount+5; i++ {
		ss = append(ss, []byte("small sample"))
	}
	ss[0] = big
	got := statSamples(ss)
	if len(got) != dictStatCount || len(got[0]) != maxBlock || string(got[1]) != "small sample" {
		t.Fatalf("%d %d", len(got), len(got[0]))
	}
	if _, err := (&Zstd{DictSize: 4096}).Train(ss[:40], FirstDictID); err != nil {
		t.Fatal(err)
	}
}
