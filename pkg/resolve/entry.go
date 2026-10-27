//ff:type feature=resolve type=model
//ff:what Entry — one GEULPROXY item: its Source, public flag, `|` fall-through and `off`
package resolve

// Entry is one item of the source list (GEUL-ORG GEULPROXY). Name is the item
// as written. Public marks geul.org and registry, which a private document
// skips. Fallthrough is set when `|` follows the item: any error of this source
// moves on, not only ErrNotFound. Off stops the list.
type Entry struct {
	Name        string
	Source      Source
	Public      bool
	Fallthrough bool
	Off         bool
}
