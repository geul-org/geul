//ff:func feature=page type=helper control=iteration dimension=1
//ff:what refs — Edge Ids as links carrying their labels
package page

// refs turns Edge Ids into links, each with the label of its record. 0 is the
// REF "none" and is left out.
func refs[T ~uint64](ids []T, labels map[uint64]string) []Ref {
	var out []Ref
	for _, id := range ids {
		if id == 0 {
			continue
		}
		out = append(out, Ref{ID: uint64(id), Label: labels[uint64(id)]})
	}
	return out
}
