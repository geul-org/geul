//ff:func feature=schema type=rule control=sequence
//ff:what checkDocument — the document-wide rules: singles, effective order, Edge Id uniqueness, CODEC targets, slot markers in TRIPLE·QUALIFIER, NO-VALUE conflicts, version list
package schema

// checkDocument applies the rules that look at more than one record.
func checkDocument(parts docParts) []Fact {
	fs := countSingles(parts.Head)
	fs = append(fs, checkEffectiveOrder(parts.Head)...)
	fs = append(fs, edgeIDDups(parts.Body)...)
	fs = append(fs, codecDecls(parts.Body)...)
	fs = append(fs, statementMarkers(parts.Body)...)
	fs = append(fs, noValueConflicts(parts.Body)...)
	if isVersionList(parts) {
		fs = append(fs, checkVersionList(parts)...)
	}
	return fs
}
