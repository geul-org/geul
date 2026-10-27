//ff:func feature=core type=encoder control=sequence
//ff:what tagWord — pack prefix 000, Kind, the 4-bit Length code and the 7-bit Codec into a TAG WORD
package core

// tagWord packs a GEUL TAG WORD `000K KLLL LCCC CCCC`: the Free/Proposal prefix
// 000, the 2-bit Kind, the 4-bit Length code and the 7-bit Codec.
func tagWord(k Kind, lenCode uint8, codec uint8) Word {
	return Word(uint16(k&3)<<11 | uint16(lenCode&0xF)<<7 | uint16(codec&0x7F))
}
