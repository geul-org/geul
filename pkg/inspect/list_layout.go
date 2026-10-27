//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what listLayout — a counted REF list: the count Value then n REFs
package inspect

// listLayout walks an F_WITH / F_GROUNDS list: the UNSIGNED-INTEGER count, then
// n REFs.
func listLayout(w *walker, name string, n int) {
	w.rec(name + ".count")
	for i := 0; i < n; i++ {
		w.rec(name)
	}
}
