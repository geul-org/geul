//ff:func feature=core type=decoder control=sequence
//ff:what TestFillPublicationTail — hash for PUBLICATION, reason for TOMBSTONE, nothing for DEPRECATION
package core

import "testing"

func TestFillPublicationTail(t *testing.T) {
	p := Publication{Codec: EdgePublication}
	fillPublicationTail(&p, append(make([]Word, 15), 0xABCD))
	if p.Hash[30] != 0xAB || p.Hash[31] != 0xCD {
		t.Fatalf("hash: %x", p.Hash)
	}
	tomb := Publication{Codec: EdgeTombstone}
	fillPublicationTail(&tomb, []Word{7})
	if tomb.Reason != 7 {
		t.Fatalf("reason: %d", tomb.Reason)
	}
	dep := Publication{Codec: EdgeDeprecation}
	fillPublicationTail(&dep, nil)
	if dep.Reason != 0 || dep.Hash != ([32]byte{}) {
		t.Fatalf("deprecation: %+v", dep)
	}
}
