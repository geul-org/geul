//ff:func feature=page type=helper control=iteration dimension=1
//ff:what wholeRecords — whether a run of at least 2 words is exactly a sequence of named Value or Reference records
package page

import "github.com/geul-org/gogeul/pkg/core"

// wholeRecords reports whether ws (at least 2 words) splits exactly into
// Value or Reference records of named codecs with nothing left over (WORD
// 0x0000 is never a TAG: Reference codec 0 is reserved and unnamed). A single raw word is
// never read as a record: one word decodes as some record too easily.
func wholeRecords(ws []core.Word) bool {
	if len(ws) < 2 {
		return false
	}
	for len(ws) > 0 {
		r, n, err := core.DecodeRecord(ws)
		if err != nil || (r.Kind != core.KindValue && r.Kind != core.KindReference) || core.CodecName(r.Kind, r.Codec) == "" {
			return false
		}
		ws = ws[n:]
	}
	return true
}
