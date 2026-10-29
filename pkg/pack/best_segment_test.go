//ff:func feature=pack type=helper control=sequence
//ff:what TestBestSegment — the window over the most frequent d-mers wins; nothing frequent scores 0
package pack

import (
	"bytes"
	"testing"
)

func TestBestSegment(t *testing.T) {
	hot := bytes.Repeat([]byte("HOT-text"), 40)
	all := append(append(bytes.Repeat([]byte{0}, 300), hot...), bytes.Repeat([]byte{1}, 300)...)
	freq := dmerFreq([][]byte{hot, hot})
	at, score := bestSegment(all, freq, 0, len(all))
	if score == 0 || at < 300 || at+coverSegment > 300+len(hot) {
		t.Fatal(at, score)
	}
	if _, score := bestSegment(all, map[uint64]uint32{}, 0, len(all)); score != 0 {
		t.Fatal(score)
	}
}
