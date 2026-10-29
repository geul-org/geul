//ff:func feature=pack type=helper control=sequence
//ff:what objBytes — a crafted object: header then body, with its OBJECTS entry
package pack

import "crypto/sha256"

// objBytes builds an object of kind holding body right after the header and
// inv (data packs) after it.
func objBytes(kind byte, body, inv []byte) ([]byte, object) {
	size := uint64(headerSize + len(body) + len(inv))
	h := objectHeader{kind: kind, size: size}
	if kind == kindData {
		h.invOff, h.invLen = uint64(headerSize+len(body)), uint64(len(inv))
	}
	b := append(encodeObjectHeader(h), body...)
	b = append(b, inv...)
	return b, object{kind: kind, size: size, hash: sha256.Sum256(b)}
}
