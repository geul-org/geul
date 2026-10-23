package resolve

const (
	// stateNotYet: the version's EFFECTIVE-FROM is after the instant.
	stateNotYet effState = iota
	// stateInEffect: the version is in effect at the instant.
	stateInEffect
	// stateExpired: the version's EFFECTIVE-UNTIL is at or before the instant.
	stateExpired
)
