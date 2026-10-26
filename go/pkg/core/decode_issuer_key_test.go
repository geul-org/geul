//ff:func feature=core type=decoder control=sequence
//ff:what TestDecodeIssuerKey — another record and a wrong width or algorithm fail
package core

import "testing"

func TestDecodeIssuerKey(t *testing.T) {
	if _, _, err := DecodeIssuerKey(EncodeNextLocalID(1)); err == nil {
		t.Fatal("other Meta: want error")
	}
	if _, _, err := DecodeIssuerKey(Record{Kind: KindMeta, Codec: MetaIssuerKey, Payload: []Word{1, 2}}); err == nil {
		t.Fatal("short: want error")
	}
	if _, _, err := DecodeIssuerKey(Record{Kind: KindMeta, Codec: MetaIssuerKey, Payload: make([]Word, 17)}); err == nil {
		t.Fatal("algorithm 0: want error")
	}
}
