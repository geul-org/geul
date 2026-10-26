//ff:func feature=cli type=helper control=selection
//ff:what logTime — a publication time in UTC, or - when unknown
package cli

import "time"

// logTime formats t as `2006-01-02 15:04:05` UTC; a zero time is `-`.
func logTime(t time.Time) string {
	switch {
	case t.IsZero():
		return "-"
	default:
		return t.UTC().Format("2006-01-02 15:04:05")
	}
}
