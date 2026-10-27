//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what noValueConflicts — E6 (Warn): a NO-VALUE TRIPLE beside a TRIPLE with a value (or UNKNOWN-VALUE) of the same Subject·Predicate·rank whose qualifiers do not tell them apart
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// noValueConflicts warns (never fails) on each TRIPLE whose Object is the
// NO-VALUE marker while a TRIPLE with the same Predicate, Subject and rank and
// the same QUALIFIER set (qualifierSets) has any other Object: "no value" and
// "a value" are then said of the same statement scope (edge/TRIPLE). Wikidata
// holds such states too, so readers settle them (rank, truthy); the checker
// only reports.
func noValueConflicts(body []rec) []Fact {
	sts := statementsOf(body)
	sets := qualifierSets(sts)
	keys := map[int]string{}
	valued := map[string]uint64{}
	for i, s := range sts {
		if s.Codec != core.EdgeTriple || len(s.Rest) != 1 {
			continue
		}
		keys[i] = core.FormatID(s.Parts[0]) + core.FormatID(s.Parts[1]) + core.FormatID(s.Rest) + "|" + sets[s.ID]
		if !(len(s.Parts[2]) == 1 && s.Parts[2][0] == core.NoValue) {
			valued[keys[i]] = s.ID
		}
	}
	var fs []Fact
	for i, s := range sts {
		k, ok := keys[i]
		if id, has := valued[k]; ok && has && len(s.Parts[2]) == 1 && s.Parts[2][0] == core.NoValue {
			fs = append(fs, factAt("E6", Warn, s.R, "Object", "같은 Subject·Predicate·rank·한정 조건에 NO-VALUE만", fmt.Sprintf("#%d NO-VALUE와 #%d 값", s.ID, id)))
		}
	}
	return fs
}
