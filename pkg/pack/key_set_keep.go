//ff:func feature=pack type=helper control=sequence
//ff:what keySet.keep — record key unless a newer layer already did
package pack

// keep stores l for k when k is not yet present.
func (s keySet) keep(k string, l Loc) error {
	if _, ok := s[k]; !ok {
		s[k] = l
	}
	return nil
}
