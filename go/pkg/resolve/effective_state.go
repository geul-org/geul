//ff:func feature=resolve type=rule control=sequence
//ff:what effectiveState — a version's standing at an instant from its EFFECTIVE-FROM and EFFECTIVE-UNTIL
package resolve

import (
	"time"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/inspect"
)

// effectiveState places doc relative to at. Without EFFECTIVE-FROM the version
// counts as in effect from the start (the provisional version list carries no
// publication time to fall back on). Without EFFECTIVE-UNTIL it never expires
// on its own; a newer version in effect supersedes it instead.
func effectiveState(doc inspect.Document, at time.Time) (effState, error) {
	from, hasFrom, err := effectiveBound(doc, core.MetaEffectiveFrom)
	if err != nil {
		return 0, err
	}
	until, hasUntil, err := effectiveBound(doc, core.MetaEffectiveUntil)
	if err != nil {
		return 0, err
	}
	if hasFrom && at.Before(from.Time) {
		return stateNotYet, nil
	}
	if hasUntil && !at.Before(until.Time) {
		return stateExpired, nil
	}
	return stateInEffect, nil
}
