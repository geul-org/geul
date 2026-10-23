//ff:func feature=schema type=decoder control=sequence
//ff:what readRec — frame one record (F2~F5); false when its length cannot be trusted
package schema

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// readRec frames the record at ws[at]: prefix (F4), Length (F2, F3) and the
// Codec extension TAG of an ISSUER codec (F5). The bool is false when the
// Length cannot be trusted; framing stops there.
func readRec(ws []core.Word, at, index int) (rec, []Fact, bool) {
	tag := ws[at]
	r := rec{Index: index, Word: at}
	var fs []Fact
	if p := tag >> 13; p == 1 {
		fs = append(fs, factAt("F4", Warn, r, "TAG WORD", "접두사 000 (Proposal)", "001 (Standard, 미배정)"))
	} else if p != 0 {
		fs = append(fs, factAt("F4", Error, r, "TAG WORD", "접두사 000", fmt.Sprintf("%03b", p)))
	}
	total, head, canonical, ok := readLength(ws, at)
	if !ok || at+total > len(ws) {
		return r, append(fs, factAt("F2", Error, r, "Length", "입력 안에 들어가는 길이", fmt.Sprintf("%d WORD (남은 입력 %d WORD)", total, len(ws)-at))), false
	}
	if !canonical {
		fs = append(fs, factAt("F3", Error, r, "Length", "가장 짧은 Length 표현", fmt.Sprintf("%d WORD를 %d WORD TAG로 표현", total, head)))
	}
	r.Total = total
	r.R = core.Record{Kind: core.Kind((tag >> 11) & 3), Codec: uint8(tag & 0x7F)}
	body := ws[at+head : at+total]
	w := core.IssuerWidth(r.R.Kind, r.R.Codec)
	if w > 0 && len(body) < w+1 {
		r.R.Payload = body
		return r, append(fs, factAt("F5", Error, r, "Codec 확장 TAG", fmt.Sprintf("Issuer %d WORD + 형식 번호 1 WORD", w), fmt.Sprintf("%d WORD", len(body)))), true
	}
	if w > 0 {
		r.R.Issuer, r.R.Format, body = body[:w], uint16(body[w]), body[w+1:]
	}
	r.R.Payload = body
	return r, fs, true
}
