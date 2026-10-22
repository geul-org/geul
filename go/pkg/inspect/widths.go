//ff:type feature=inspect type=model
//ff:what Widths — the Edge Id, Variable ID and IMPORT ID widths in words (Meta ID-WIDTH)
package inspect

// Widths are the three ID widths in words that Meta ID-WIDTH declares (GEUL
// t002 OVERVIEW ID-WIDTH): Edge Id, Variable ID and IMPORT ID.
type Widths struct {
	Edge   int
	Var    int
	Import int
}
