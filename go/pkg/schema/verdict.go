//ff:type feature=schema type=model
//ff:what Verdict — Pass, Fail or Review for a whole document
package schema

// Verdict is the outcome for a document: Fail with any Error, Review with only
// Warns, Pass with no Fact.
type Verdict int
