//ff:func feature=schema type=rule control=sequence
//ff:what checkEffectiveOrder — M5: EFFECTIVE-FROM comes before EFFECTIVE-UNTIL
package schema

import "github.com/geul-org/gogeul/pkg/core"

// checkEffectiveOrder applies M5 when both bounds are present and valid.
func checkEffectiveOrder(head []rec) []Fact {
	from, _, okFrom := metaInstant(head, core.MetaEffectiveFrom)
	until, r, okUntil := metaInstant(head, core.MetaEffectiveUntil)
	if okFrom && okUntil && !from.Before(until) {
		return []Fact{factAt("M5", Error, r, "", "EFFECTIVE-FROM < EFFECTIVE-UNTIL", from.Format("2006-01-02")+" ≥ "+until.Format("2006-01-02"))}
	}
	return nil
}
