package schema

const (
	// Error is a spec violation.
	Error Severity = iota
	// Warn is a reserved or unassigned code that a newer spec may define.
	Warn
)
