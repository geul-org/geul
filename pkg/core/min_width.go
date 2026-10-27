//ff:func feature=core type=helper control=sequence
//ff:what MinWidth — the narrowest 1/2/4-word field width that holds max
package core

// MinWidth returns the narrowest declarable width (1, 2 or 4 words) whose
// unsigned range holds max. Fixed-width fields (the DOC-* length, the TANGEUL
// SPAN width) are sized this way, so a writer fits them to its own content.
func MinWidth(max uint64) int {
	if max <= 0xFFFF {
		return 1
	}
	if max <= 0xFFFFFFFF {
		return 2
	}
	return 4
}
