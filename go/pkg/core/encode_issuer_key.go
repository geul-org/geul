//ff:func feature=core type=encoder control=sequence
//ff:what EncodeIssuerKey — an ISSUER-KEY Meta: signature algorithm ID and public key
package core

import "fmt"

// EncodeIssuerKey builds ISSUER-KEY (meta/verification/ISSUER-KEY): the
// algorithm ID word then the public key. Only Ed25519 (32-byte key) is
// defined.
func EncodeIssuerKey(alg Word, pub []byte) (Record, error) {
	if alg != SigEd25519 || len(pub) != 32 {
		return Record{}, fmt.Errorf("geul/core: ISSUER-KEY needs Ed25519 (0x0001) and a 32-byte key, got 0x%04X and %d bytes", alg, len(pub))
	}
	ws, _ := ReadWords(pub)
	return Record{Kind: KindMeta, Codec: MetaIssuerKey, Payload: append([]Word{alg}, ws...)}, nil
}
