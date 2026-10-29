//ff:type feature=pack type=model
//ff:what treeWalk — collects the file paths under an archive's `.well-known/geul/`
package pack

// treeWalk gathers slash paths relative to root.
type treeWalk struct {
	root  string
	paths []string
}
