//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what writer.records — the puts that get their own record, in data pack order
package pack

import "sort"

// records lists the puts to store (shared h/ puts ride on another record).
func (w *writer) records() []*put {
	var out []*put
	for _, p := range w.puts {
		if p.via == nil {
			out = append(out, p)
		}
	}
	sort.Slice(out, func(i, j int) bool { return recordOrder(out[i], w.inum) < recordOrder(out[j], w.inum) })
	return out
}
