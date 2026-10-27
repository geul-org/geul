//ff:type feature=schema type=model
//ff:what Report — every Fact found in a document and the resulting Verdict
package schema

// Report is the result of Check: all Facts in document order of discovery and
// the Verdict they add up to.
type Report struct {
	Facts   []Fact
	Verdict Verdict
}
