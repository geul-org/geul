//ff:func feature=core type=encoder control=sequence
//ff:what issuerOfDigest — a 32-byte digest as a 16-word Issuer ID, refused when below 2^64 (not canonical as 16 words)
package core

import "errors"

// issuerOfDigest reads digest as 16 words high byte first; a value below
// 2^64 would not be canonical at that width, so that key makes no Issuer.
func issuerOfDigest(digest []byte) ([]Word, error) {
	ws, _ := ReadWords(digest)
	if !IssuerCanonical(ws) {
		return nil, errors.New("geul/core: key hashes below 2^64; it makes no 16-word Issuer")
	}
	return ws, nil
}
