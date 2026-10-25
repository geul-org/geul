//ff:func feature=core type=helper control=selection
//ff:what IssuerCanonical — whether an Issuer is written in the shortest of 1, 2, 4 and 16 words that holds its value
package core

import "slices"

// IssuerCanonical reports whether issuer is an Issuer ID in canonical form
// (GEUL t002 ISSUER): 1, 2, 4 or 16 words, and the shortest of those that holds
// its value. An Issuer ID is its integer value, so 0x0005 and 0x00000005 are
// the same Issuer; a width grows only when the value does not fit the
// narrower one (2 WORD from 0x00010000, 4 WORD from 2^32, 16 WORD from 2^64).
func IssuerCanonical(issuer []Word) bool {
	switch len(issuer) {
	case 1:
		return true
	case 2:
		return issuer[0] != 0
	case 4:
		return issuer[0]|issuer[1] != 0
	case 16:
		return slices.ContainsFunc(issuer[:12], func(w Word) bool { return w != 0 })
	default:
		return false
	}
}
