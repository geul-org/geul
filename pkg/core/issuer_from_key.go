//ff:func feature=core type=encoder control=sequence
//ff:what IssuerFromKey — the 16-word Issuer ID of a public key: SHA-256("GEUL-ISSUER-v1" ‖ algorithm ID ‖ key)
package core

import "crypto/sha256"

// IssuerFromKey computes a self-issued Issuer ID (OVERVIEW 16 WORD Issuer 생성
// 규칙): SHA-256 of "GEUL-ISSUER-v1", the algorithm ID (2 bytes, high byte
// first) and the public key, written high byte first into 16 words. A result
// below 2^64 would not be canonical as 16 words and is refused: that key makes
// no Issuer.
func IssuerFromKey(alg Word, pub []byte) ([]Word, error) {
	h := sha256.New()
	h.Write([]byte("GEUL-ISSUER-v1"))
	h.Write([]byte{byte(alg >> 8), byte(alg)})
	h.Write(pub)
	return issuerOfDigest(h.Sum(nil))
}
