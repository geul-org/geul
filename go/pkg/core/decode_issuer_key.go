//ff:func feature=core type=decoder control=selection
//ff:what DecodeIssuerKey — the algorithm ID and public key of an ISSUER-KEY Meta
package core

import "fmt"

// DecodeIssuerKey reads ISSUER-KEY: Ed25519 (0x0001) and 16 words of key.
// Another record, another algorithm or another key width is an error.
func DecodeIssuerKey(r Record) (Word, []byte, error) {
	switch {
	case r.Kind != KindMeta || r.Codec != MetaIssuerKey:
		return 0, nil, fmt.Errorf("geul/core: expected Meta ISSUER-KEY, got kind %d codec 0x%02X", r.Kind, r.Codec)
	case len(r.Payload) != 17 || r.Payload[0] != SigEd25519:
		return 0, nil, fmt.Errorf("geul/core: ISSUER-KEY must be Ed25519 with a 16-word key (%d words)", len(r.Payload))
	}
	return r.Payload[0], PutWords(r.Payload[1:]), nil
}
