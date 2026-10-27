//ff:func feature=page type=helper control=selection
//ff:what tagBits — a TAG WORD's bit fields: prefix, Kind, Length, Codec and, for a Reference, the Quantifier
package page

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// tagBits spells out TAG WORD `000K KLLL LCCC CCCC`: the prefix, the Kind, the
// Length code with the word count it gives (or the extension it calls for) and
// the Codec with its spec name, else fallback (a profile record name), else
// hex. A Reference splits the codec into CCCC and the Quantifier QQQ.
func tagBits(w core.Word, fallback string) string {
	k, l, codec := core.Kind((w>>11)&3), uint8((w>>7)&0xF), uint8(w&0x7F)
	length := fmt.Sprintf("%d WORDs", l+1)
	switch l {
	case 0:
		length = "1 WORD"
	case 0xE:
		length = "Length 확장 1 WORD"
	case 0xF:
		length = "Length 확장 2 WORD"
	}
	name := core.CodecName(k, codec)
	if name == "" {
		name = fallback
	}
	if name == "" {
		name = fmt.Sprintf("0x%02X", codec)
	}
	head := fmt.Sprintf("%03b Proposal\n%02b %s\n%04b %s", uint16(w)>>13, uint8(k), k, l, length)
	if k == core.KindReference {
		return fmt.Sprintf("%s\n%04b %s\n%03b %s", head, codec>>3, name, codec&7, core.QuantifierName(codec))
	}
	return fmt.Sprintf("%s\n%07b %s", head, codec, name)
}
