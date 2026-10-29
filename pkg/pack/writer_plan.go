//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what writer.plan — the changed puts (bytes differ from the previous generation) and the real deletions
package pack

import (
	"slices"
)

// plan reads every Put path once, keeps those whose bytes changed, and lists
// deletions of paths the previous generation has (Full deletes all others).
func (w *writer) plan() error {
	paths := slices.Clone(w.b.Put)
	slices.Sort(paths)
	paths = slices.Compact(paths)
	w.sums = make(map[string][32]byte, len(paths))
	for _, p := range paths {
		pu, err := w.planPut(p)
		if err != nil {
			return err
		}
		if pu != nil {
			w.puts = append(w.puts, pu)
		}
	}
	w.planDeletes()
	return w.share()
}
