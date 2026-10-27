//ff:func feature=cli type=helper control=selection
//ff:what logTime — a publication time in UTC, or - when unknown
package cli

import "time"

// logTime formats t as RFC 3339 UTC (`2006-01-02T15:04:05Z`, the zone shown
// so the time is not read as local); a zero time is `-`.
func logTime(t time.Time) string {
	switch {
	case t.IsZero():
		return "-"
	default:
		return t.UTC().Format(time.RFC3339)
	}
}
