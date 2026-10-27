//ff:func feature=schema type=rule control=sequence
//ff:what Check — every layer-1 rule, then the given Codecs, over one .geul byte string
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// Check examines b as one .geul document: framing, structure, per-Kind rules,
// document-wide rules and the version list rules (layer 1, README), then every
// Codec in opt (layer 2). It keeps going after a violation as long as record
// boundaries can be trusted.
func Check(b []byte, opt Options) Report {
	var fs []Fact
	if len(b)%2 != 0 {
		fs = append(fs, Fact{Rule: "F1", Severity: Error, Where: Where{Word: len(b) / 2}, Expected: "짝수 바이트", Actual: fmt.Sprintf("%d 바이트", len(b))})
		b = b[:len(b)-1]
	}
	ws, _ := core.ReadWords(b)
	recs, f := frame(ws)
	fs = append(fs, f...)
	parts, f := structure(recs, len(ws))
	fs = append(fs, f...)
	fs = append(fs, checkRecords(parts)...)
	fs = append(fs, checkDocument(parts)...)
	fs = append(fs, runCodecs(recs, opt.Codecs)...)
	if opt.Strict {
		fs = strictify(fs)
	}
	return Report{Facts: fs, Verdict: verdictOf(fs)}
}
