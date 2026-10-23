package schema

const (
	// Pass means no Fact was found.
	Pass Verdict = iota
	// Fail means at least one Error.
	Fail
	// Review means only Warns: possibly valid under a newer spec.
	Review
)
