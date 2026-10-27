//ff:func feature=page type=builder control=iteration dimension=1
//ff:what summaryOf — a record's field texts joined for its one-line list entry
package page

import "strings"

// summaryOf joins the texts of fs (the row's fields, in order) with spaces, so
// a list entry shows what the record says (`en Q1 "Universe"`) and Edges of the
// same codec can be told apart. Fields without text are left out.
func summaryOf(fs []Field) string {
	var parts []string
	for _, f := range fs {
		if f.Text != "" {
			parts = append(parts, f.Text)
		}
	}
	return strings.Join(parts, " ")
}
