//ff:type feature=core type=model
//ff:what EnumDecl — an ENUM Edge: Edge Id, the FIELD it belongs to, the code words
package core

// EnumDecl is an ENUM Edge (GEUL t002 docs/edge/ENUM): one code of the raw
// FIELD Field, as many words as that field's width.
type EnumDecl struct {
	EdgeID uint64
	Field  uint64
	Code   []Word
}
