//ff:func feature=core type=decoder control=sequence
//ff:what DecodeEnumDecl — an ENUM Edge as EnumDecl
package core

import "fmt"

// DecodeEnumDecl reads an ENUM Edge (GEUL t002 docs/edge/ENUM): the LOCAL Edge
// Id, a LOCAL Reference to its FIELD and the code words to the end. Whether the
// code width matches the field is the codec spec's check.
func DecodeEnumDecl(r Record) (EnumDecl, error) {
	if r.Kind != KindEdge || r.Codec != EdgeEnum {
		return EnumDecl{}, fmt.Errorf("geul/core: not an ENUM Edge (kind %d codec 0x%02X)", r.Kind, r.Codec)
	}
	id, pos, err := takeLocal(r.Payload)
	if err != nil {
		return EnumDecl{}, err
	}
	f, n, err := takeLocal(r.Payload[pos:])
	if err != nil {
		return EnumDecl{}, err
	}
	if pos+n == len(r.Payload) {
		return EnumDecl{}, fmt.Errorf("geul/core: ENUM without a code")
	}
	return EnumDecl{EdgeID: id, Field: f, Code: r.Payload[pos+n:]}, nil
}
