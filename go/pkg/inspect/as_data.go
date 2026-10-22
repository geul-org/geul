//ff:func feature=inspect type=helper control=sequence
//ff:what asData — a parse result as Entry.Data: the value when ok, nil otherwise
package inspect

// asData returns v when ok and nil otherwise, so a failed parse leaves
// Entry.Data empty rather than holding a zero value.
func asData[T any](v T, ok bool) any {
	if !ok {
		return nil
	}
	return v
}
