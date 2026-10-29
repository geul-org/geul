//ff:type feature=resolve type=resolver
//ff:what dirStore — Store over a plain directory (`<Root>/.well-known/geul/`), one file per path
package resolve

// dirStore is the directory layout every Dir used before stores existed.
type dirStore struct {
	root string
}
