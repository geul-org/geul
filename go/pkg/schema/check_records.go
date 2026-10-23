//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what checkRecords — the per-record rules over HEAD and BODY
package schema

// checkRecords applies the per-record rules to every HEAD and BODY record,
// plus the IMPORT ID duplicate rule the Reference rules depend on.
func checkRecords(parts docParts) []Fact {
	imports, fs := importsOf(parts.Head)
	for _, r := range append(append([]rec{}, parts.Head...), parts.Body...) {
		fs = append(fs, checkRecord(r, imports)...)
	}
	return fs
}
