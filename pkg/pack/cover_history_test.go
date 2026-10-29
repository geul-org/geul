//ff:func feature=pack type=encoder control=iteration dimension=1
//ff:what TestCoverHistory — at most size bytes of shared content, deterministic, best last; nothing shared or too little gives none
package pack

import (
	"bytes"
	"testing"
)

func TestCoverHistory(t *testing.T) {
	_, get := synth(64)
	var samples [][]byte
	for i := 0; i < 64; i++ {
		b, _ := get(synthPath(i))
		samples = append(samples, b)
	}
	h := coverHistory(samples, 2048)
	if len(h) == 0 || len(h) > 2048 || len(h)%coverSegment != 0 || !bytes.Equal(h, coverHistory(samples, 2048)) {
		t.Fatal(len(h))
	}
	if !bytes.Contains(h[len(h)-coverSegment:], []byte("source=\"wikidata dump 2026-10-19\"")) {
		t.Fatalf("%q", h[len(h)-coverSegment:])
	}
	if h := coverHistory([][]byte{bytes.Repeat([]byte{1}, 600)}, 2048); len(h) != coverSegment {
		t.Fatal("the second epoch scores 0 once the first pick zeroed its d-mers", len(h))
	}
	if h := coverHistory([][]byte{[]byte("short")}, 2048); h != nil {
		t.Fatal("too little")
	}
}
