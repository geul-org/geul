//ff:func feature=page type=builder control=sequence
//ff:what TestSummaryOf — field texts joined by spaces in order; empty texts dropped
package page

import "testing"

func TestSummaryOf(t *testing.T) {
	if got := summaryOf([]Field{{Key: "language", Text: "en"}, {Key: "subject", Text: "Q1"}, {Key: "x"}, {Key: "text", Text: `"Universe"`}}); got != `en Q1 "Universe"` {
		t.Fatalf("got %q", got)
	}
	if got := summaryOf(nil); got != "" {
		t.Fatalf("empty: %q", got)
	}
}
